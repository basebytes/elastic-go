package transport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Interface interface {
	Perform(*http.Request) (*http.Response, error)
}

type Client struct {
	sync.Mutex
	servers    []string
	username   string
	password   string
	transport  http.RoundTripper
	maxRetries int

	poolFunc func([]*Connection, Selector) ConnectionPool
	pool     ConnectionPool
	selector Selector
}

func New(cfg *Config) (*Client, error) {
	var conns []*Connection
	for _, server := range cfg.Servers {
		if u, err := url.Parse(server); err == nil {
			if cfg.Username != "" {
				u.User = url.UserPassword(cfg.Username, cfg.Password)
			}
			conns = append(conns, &Connection{URL: u})
		} else {
			return nil, err
		}
	}

	if cfg.MaxRetries <= 1 {
		cfg.MaxRetries = 1
	}
	if cfg.Transport == nil {
		cfg.Transport = http.DefaultTransport
	}

	client := Client{
		servers:    cfg.Servers,
		transport:  cfg.Transport,
		username:   cfg.Username,
		password:   cfg.Password,
		maxRetries: cfg.MaxRetries,
	}
	if client.poolFunc != nil {
		client.pool = client.poolFunc(conns, client.selector)
	} else {
		client.pool, _ = NewConnectionPool(conns, client.selector)
	}
	return &client, nil
}

func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	var (
		res *http.Response
		err error
	)
	if req.Body != nil && req.Body != http.NoBody && req.GetBody == nil {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(req.Body)
		req.GetBody = func() (io.ReadCloser, error) {
			r := buf
			return ioutil.NopCloser(&r), nil
		}
		if req.Body, err = req.GetBody(); err != nil {
			return nil, fmt.Errorf("cannot get request body: %s", err)
		}
	}

	for i := 0; i <= c.maxRetries; i++ {
		var (
			conn        *Connection
			shouldRetry bool
			//shouldCloseBody bool
		)
		c.Lock()
		conn, err = c.pool.Next()
		c.Unlock()
		if err != nil {
			//if c.logger != nil {
			//	c.logRoundTrip(req, nil, err, time.Time{}, time.Duration(0))
			//}
			return nil, fmt.Errorf("cannot get connection: %s", err)
		}
		c.setReqURL(conn.URL, req)
		c.setReqAuth(conn.URL, req)
		if i > 1 && req.Body != nil && req.Body != http.NoBody {
			body, err := req.GetBody()
			if err != nil {
				return nil, fmt.Errorf("cannot get request body: %s", err)
			}
			req.Body = body
		}
		//start := time.Now().UTC()
		res, err = c.transport.RoundTrip(req)
		//dur := time.Since(start)
		if err != nil {
			// Report the connection as unsuccessful
			c.Lock()
			c.pool.OnFailure(conn)
			c.Unlock()
			if err == io.EOF {
				shouldRetry = true
			}
			if err, ok := err.(net.Error); ok {
				shouldRetry = !err.Timeout()
			}
		} else {
			c.Lock()
			c.pool.OnSuccess(conn)
			c.Unlock()
		}
		if !shouldRetry {
			break
		}
	}
	return res, err
}

func (c *Client) setReqURL(u *url.URL, req *http.Request) *http.Request {
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	if u.Path != "" {
		var b strings.Builder
		b.Grow(len(u.Path) + len(req.URL.Path))
		b.WriteString(u.Path)
		b.WriteString(req.URL.Path)
		req.URL.Path = b.String()
	}

	return req
}

func (c *Client) setReqAuth(u *url.URL, req *http.Request) *http.Request {
	if _, ok := req.Header["Authorization"]; !ok {
		if u.User != nil {
			pwd, _ := u.User.Password()
			req.SetBasicAuth(u.User.Username(), pwd)
			return req
		}
	}
	return req
}
