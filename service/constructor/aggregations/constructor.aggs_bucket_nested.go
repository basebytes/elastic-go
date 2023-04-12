package aggregations

func NewNested() Nested {
	return func(path string, o ...func(*NestedParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"nested": map[string]interface{}{
					"path": path,
				},
			}
		}
		b := &NestedParam{path: path, param: make(map[string]interface{})}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Nested func(path string, o ...func(*NestedParam)) map[string]interface{}

func (n Nested) WithChildAgg(agg map[string]interface{}) func(*NestedParam) {
	return func(p *NestedParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type NestedParam struct {
	path  string
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (n *NestedParam) Build() map[string]interface{} {
	res := map[string]interface{}{
		"nested": map[string]interface{}{
			"path": n.path,
		},
	}
	if len(n.aggs) > 0 {
		res["aggs"] = n.aggs

	}
	return res
}
