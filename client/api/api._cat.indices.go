package api

import "strings"

func newCatIndices(transport Transport) CatIndices {
	return func(o ...func(*CatIndicesRequest)) (*Response, error) {
		var r = CatIndicesRequest{}
		r.uris= r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type CatIndices func(o ...func(*CatIndicesRequest)) (*Response, error)

func (f *CatIndices) WithIndex(v ...string) func(*CatIndicesRequest) {
	return func(r *CatIndicesRequest) {
		r.Index = strings.Join(v, ",")
	}
}

func (f *CatIndices) WithColumns(v ...string) func(*CatIndicesRequest) {
	return func(r *CatIndicesRequest) {
		r.H = strings.Join(v, ",")
		r.AddParam("h", r.H)
	}
}

func (f *CatIndices) WithFormat(v string) func(*CatIndicesRequest) {
	return func(r *CatIndicesRequest) {
		r.Format = strings.TrimSpace(v)
		r.AddParam("format", r.Format)
	}
}

type CatIndicesRequest struct {
	BaseRequest
	H      string
	Format string
}

func (r *CatIndicesRequest)getUris() []string{
	var uris = []string{"_cat", "indices"}
	if r.Index != "" {
		uris = append(uris, r.Index)
	}
	return uris
}