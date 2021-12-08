package api

import "strings"

func newDocumentGetFunc(transport Transport) Get {
	return func(index,_id string, o ...func(*GetRequest)) (*Response, error) {
		r := GetRequest{BaseRequest: BaseRequest{Index: index},DocumentID: _id}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type Get func(index,_id string,o ...func(*GetRequest))(*Response,error)

func (g Get) WithIncludes(v ...string) func(*GetRequest) {
	return func(p *GetRequest) {
		if len(v)>0{
			p.AddParam("_source_includes", strings.Join(v,","))
		}
	}
}

func (g Get) WithExcludes(v ...string) func(*GetRequest) {
	return func(p *GetRequest) {
		if len(v)>0{
			p.AddParam("_source_excludes",  strings.Join(v,","))
		}
	}
}
type GetRequest struct {
	BaseRequest
	DocumentID string
}

func (g *GetRequest) getUris() []string {
	var uris []string
	uris = append(uris, g.Index, "_source")
	if g.DocumentID != "" {
		uris = append(uris, g.DocumentID)
	}
	return uris
}