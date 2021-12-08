package entity

import "strings"

//alias
type IndexAlias struct {
	EsError	`mapstructure:",squash"`
	Other  map[string]*IndexAliasItem `json:"result,omitempty" mapstructure:",remain"`
}

type IndexAliasItem struct {
	Aliases map[string]interface{}	`json:"aliases,omitempty"`
}

func (a *IndexAlias) GetIndexByAliasName(name string) []string {
	if len(a.Other) == 0 {
		return nil
	} else {
		var indices []string
		name = strings.ToLower(name)
		for k, v := range a.Other {
			if _, OK := v.Aliases[name]; OK {
				indices = append(indices, k)
			}
		}
		return indices
	}
}



