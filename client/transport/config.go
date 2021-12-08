package transport

import (
	"net/http"
)

type Config struct {
	Servers []string
	MaxRetries int
	Transport http.RoundTripper
	ConnectionPoolFunc func([]*Connection, Selector) ConnectionPool
	pool      ConnectionPool
	selector  Selector
}