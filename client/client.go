package client
import (
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/elastic-go/client/transport"

	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	defaultURL = "http://localhost:9200"
)

type Config struct {
	Servers []string
	MaxRetries int
	Transport http.RoundTripper
	ConnectionPoolFunc func([]*transport.Connection, transport.Selector) transport.ConnectionPool
	pool      transport.ConnectionPool
	selector  transport.Selector
}

type Client struct {
	*api.API
	Transport transport.Interface
}

func NewDefaultClient() (*Client, error) {
	return NewClient(&Config{})
}

func NewClient(cfg *Config) (*Client, error) {
	var addrs []string
	if len(cfg.Servers)==0{
		addrs = addrsFromEnvironment()
	}else{
		addrs = append(addrs, cfg.Servers...)
	}
	if len(addrs)==0{
		addrs=append(addrs, defaultURL)
	}
	if cfg.MaxRetries<=1{
		cfg.MaxRetries=1
	}

	tp, err := transport.New(&transport.Config{
		Servers:            addrs,
		MaxRetries:         cfg.MaxRetries,
		Transport:          cfg.Transport,
		ConnectionPoolFunc: cfg.ConnectionPoolFunc,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}
	client := &Client{
		API:       api.New(tp),
		Transport: tp,
	}
	return client, nil
}

func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	return c.Transport.Perform(req)
}

func addrsFromEnvironment() []string {
	var addrs []string
	if envURLs, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURLs != "" {
		list := strings.Split(envURLs, ",")
		for _, u := range list {
			addrs = append(addrs, strings.TrimSpace(u))
		}
	}

	return addrs
}