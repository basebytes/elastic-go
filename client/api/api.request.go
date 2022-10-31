package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	headerContentType         = "Content-Type"
	headerContentTypeProtobuf = "application/x-protobuf"
	headerContentTypeJSONUtf8 = "application/json;charset=utf-8"
	defaultMethod             = http.MethodGet
)

type Request interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
}

type BaseRequest struct {
	Index  string
	Body   *[]byte
	params map[string]string

	ctx    context.Context
	uris   func() []string
	method func() string
}

func (b *BaseRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	httpRequest, err := http.NewRequest(b.getMethod(), strings.Join(b.getUris(), "/"), nil)
	b.addParams(httpRequest)

	//Accept+";compatible-with=7"
	if err != nil {
		return nil, err
	} else if ctx != nil {
		httpRequest = httpRequest.WithContext(ctx)
	} else if b.Body != nil {
		httpRequest.Body = ioutil.NopCloser(bytes.NewReader(*b.Body))
		httpRequest.Header.Add(headerContentType, headerContentTypeJSONUtf8)
	}

	res, err := transport.Perform(httpRequest)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	var rspData []byte
	if res.Header.Get("Content-Encoding") == "gzip" {
		r, err := gzip.NewReader(res.Body)
		if err == nil {
			rspData, err = ioutil.ReadAll(r)
		}
	} else {
		rspData, err = ioutil.ReadAll(res.Body)
	}
	if err != nil {
		return nil, err
	}
	return &Response{Response: res, rspData: rspData, req: httpRequest}, err
}

func (b *BaseRequest) addParams(req *http.Request) {
	if b.params != nil {
		q := req.URL.Query()
		for k, v := range b.params {
			if k == "method" {
				continue
			}
			q.Set(k, v)
			//req.Param(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	//if b.Body != nil {
	//	req.Body(*b.Body)
	//	req.Header(headerContentType, headerContentTypeJSONUtf8)
	//}
}

func (b *BaseRequest) AddParam(key, value string) {
	if b.params == nil {
		b.params = make(map[string]string)
	}
	b.params[strings.ToLower(key)] = value
}

func (b *BaseRequest) getMethod() string {
	if b.method == nil {
		return MethodGetFunc()
	}
	return b.method()
}

func (b *BaseRequest) getUris() []string {
	if b.uris == nil {
		return []string{b.Index}
	}
	return b.uris()
}
