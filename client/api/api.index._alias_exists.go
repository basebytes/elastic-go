package api

import (
	"strings"
)

func newAliasExistsFunc(transport Transport) AliasExists {
	return func(alias []string, o ...func(request *AliasExistsRequest)) (*Response, error) {
		var r = AliasExistsRequest{
				Alias: alias,
				BaseRequest:BaseRequest{
					method: MethodHeadFunc,
				},
			}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type AliasExists func(alias []string, o ...func(request *AliasExistsRequest)) (*Response, error)

func (f AliasExists) WithIndex(v ...string) func(*AliasExistsRequest) {
	return func(r *AliasExistsRequest) {
		r.Index = strings.Join(v, ",")
	}
}

type AliasExistsRequest struct {
	BaseRequest
	Alias []string
}

func (r *AliasExistsRequest) getUris() []string {
	var uris []string
	if r.Index != "" {
		uris = append(uris, r.Index)
	}
	uris = append(uris, "_alias")
	uris = append(uris, strings.Join(r.Alias, ","))
	return uris
}
