package api

import (
	"strings"
)

func newIndexDelete(transport Transport) IndexDelete {
	return func(index []string, o ...func(*IndexDeleteRequest)) (*Response, error) {
		var r = IndexDeleteRequest{BaseRequest{
			Index: strings.Join(index, ","),
			method: MethodDeleteFunc,
			},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type IndexDelete func(index []string, o ...func(*IndexDeleteRequest)) (*Response, error)

type IndexDeleteRequest struct {
	BaseRequest
}