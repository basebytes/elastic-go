package constructor

func newNestedAgg() NestedAgg {
	return func(path string, o ...func(*NestedAggParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"nested": map[string]interface{}{
					"path": path,
				},
			}
		}
		b := &NestedAggParam{path: path, param: make(map[string]interface{})}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type NestedAgg func(path string, o ...func(*NestedAggParam)) map[string]interface{}

func (n NestedAgg) WithChildAgg(agg map[string]interface{}) func(*NestedAggParam) {
	return func(p *NestedAggParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type NestedAggParam struct {
	path  string
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (n *NestedAggParam) Build() map[string]interface{} {
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
