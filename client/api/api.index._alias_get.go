package api

import "strings"

func newAliasGetFunc(transport Transport) AliasGet {
	return func(o ...func(*AliasGetRequest)) (*Response, error) {
		var r = AliasGetRequest{}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type AliasGet func(o ...func(*AliasGetRequest)) (*Response, error)

func (f AliasGet) WithIndex(v ...string) func(*AliasGetRequest) {
	return func(r *AliasGetRequest) {
		if len(v)>0{
			r.Index = strings.Join(v, ",")
		}
	}
}

func (f AliasGet) WithAlias(v ...string) func(*AliasGetRequest) {
	return func(r *AliasGetRequest) {
		if len(v)>0{
			r.Alias = v
		}
	}
}

type AliasGetRequest struct {
	BaseRequest
	Alias []string
}

func (r *AliasGetRequest) getUris() []string {
	var uris []string
	if r.Index != "" {
		uris = append(uris, r.Index)
	}
	uris = append(uris, "_alias")
	uris = append(uris, strings.Join(r.Alias, ","))
	return uris
}
