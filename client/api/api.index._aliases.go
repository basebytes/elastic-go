package api

import (
	"fmt"
	"github.com/basebytes/tools"
)

func newAliasesFunc(transport Transport) Aliases {
	return func(actions []*AliasAction, o ...func(*AliasesRequest)) (*Response, error) {
		if len(actions) == 0 {
			return nil, fmt.Errorf("missing aliases action")
		}
		act := map[string][]*AliasAction{
			"actions": actions,
		}

		body := tools.EncodeBytes(act)
		var r = AliasesRequest{
			BaseRequest: BaseRequest{
				Body:   &body,
				method: MethodPostFunc,
			},
		}
		r.uris = r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type Aliases func(actions []*AliasAction, o ...func(*AliasesRequest)) (*Response, error)

type AliasesRequest struct {
	BaseRequest
}

func (r *AliasesRequest) getUris() []string {
	return []string{"_aliases"}
}

type AliasAction struct {
	Add         *ActionParam `json:"add,omitempty"`
	Remove      *ActionParam `json:"remove,omitempty"`
	RemoveIndex *ActionParam `json:"remove_index,omitempty"`
}

type ActionParam struct {
	Index         string                 `json:"index,omitempty"`
	Indices       []string               `json:"indices,omitempty"`
	Alias         string                 `json:"alias,omitempty"`
	Aliases       []string               `json:"aliases,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	IsHidden      bool                   `json:"is_hidden,omitempty"`
	MustExist     bool                   `json:"must_exist,omitempty"`
	IsWriteIndex  bool                   `json:"is_write_index,omitempty"`
	Routing       string                 `json:"routing,omitempty"`
	IndexRouting  string                 `json:"index_routing,omitempty"`
	SearchRouting string                 `json:"search_routing,omitempty"`
}
