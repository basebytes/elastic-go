package aggregations

func NewReverseNested() ReverseNested {
	return func(o ...func(*ReverseNestedParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"reverse_nested": struct{}{},
			}
		}
		b := &ReverseNestedParam{}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type ReverseNested func(o ...func(*ReverseNestedParam)) map[string]interface{}

func (r ReverseNested) WithPath(path string) func(*ReverseNestedParam) {
	return func(p *ReverseNestedParam) {
		if path != "" {
			p.path = path
		}
	}
}

func (r ReverseNested) WithChildAgg(agg map[string]interface{}) func(*ReverseNestedParam) {
	return func(p *ReverseNestedParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type ReverseNestedParam struct {
	name string
	path string
	aggs map[string]interface{}
}

func (r *ReverseNestedParam) Build() map[string]interface{} {
	res := make(map[string]interface{})
	if r.path == "" {
		res["reverse_nested"] = struct{}{}
	} else {
		res["reverse_nested"] = map[string]interface{}{
			"path": r.path,
		}
	}
	if len(r.aggs) > 0 {
		res["aggs"] = r.aggs
	}
	return res
}
