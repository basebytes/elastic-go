package helper

import (
	"github.com/basebytes/elastic-go/client/api"
)

func NewAddAction(indices, aliases []string, isWriteIndex bool) *api.AliasAction {
	return &api.AliasAction{
		Add: getActionParam(indices, aliases, isWriteIndex),
	}

}

func NewRemoveAction(indices, aliases []string, isWriteIndex bool) *api.AliasAction {
	return &api.AliasAction{
		Remove: getActionParam(indices, aliases, isWriteIndex),
	}
}

func NewRemoveIndexAction(indices, aliases []string, isWriteIndex bool) *api.AliasAction {
	return &api.AliasAction{
		RemoveIndex: getActionParam(indices, aliases, isWriteIndex),
	}
}

func getActionParam(indices, aliases []string, isWriteIndex bool) *api.ActionParam {
	p := &api.ActionParam{}
	if len(indices) == 1 {
		p.Index = indices[0]
	} else {
		p.Indices = indices
	}
	if len(aliases) == 1 {
		p.Alias = aliases[0]
	} else {
		p.Aliases = aliases
	}
	if isWriteIndex {
		p.IsWriteIndex = isWriteIndex
	}
	return p
}

func MergeAggs(aggs ...map[string]interface{}) map[string]interface{} {
	if len(aggs) == 0 || aggs[0] == nil || len(aggs[0]) == 0 {
		return nil
	}
	if len(aggs) == 1 {
		return aggs[0]
	}
	res := make(map[string]interface{})
	for _, agg := range aggs {
		for k, v := range agg {
			res[k] = v
		}
	}
	return res
}
