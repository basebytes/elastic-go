package api

import "net/http"

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
}
