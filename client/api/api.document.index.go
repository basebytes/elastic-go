package api

import (
	"net/http"
)

func newDocumentIndexFunc(transport Transport) DocumentIndex {
	return func(index string, doc Doc, o ...func(*DocumentIndexRequest)) (*Response, error) {
		body:=doc.Content()
		r := DocumentIndexRequest{BaseRequest: BaseRequest{Index: index, Body: &body}}
		r.uris=r.getUris
		r.method=r.getMethod
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type DocumentIndex func(index string, doc Doc, o ...func(*DocumentIndexRequest)) (*Response, error)

func (f DocumentIndex) WithDocumentID(v string) func(*DocumentIndexRequest) {
	return func(r *DocumentIndexRequest) {
		r.DocumentID = v
	}
}

func (f DocumentIndex) WithOpType(v string) func(*DocumentIndexRequest) {
	return func(r *DocumentIndexRequest) {
		r.OpType = v
		r.AddParam("op_type", v)
	}
}

type DocumentIndexRequest struct {
	BaseRequest
	DocumentID string
	OpType     string
}

func (r *DocumentIndexRequest) getMethod() string {
	if r.DocumentID != "" {
		return http.MethodPut
	} else {
		return http.MethodPost
	}
}

func (r *DocumentIndexRequest) getUris() []string {
	var uris []string
	uris = append(uris, r.Index, "_doc")
	if r.DocumentID != "" {
		uris = append(uris, r.DocumentID)
	}
	return uris
}
