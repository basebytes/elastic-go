package constructor

func newReverseNestedAgg() ReverseNestedAgg {
	return func(o ...func(*ReverseNestedAggParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"reverse_nested": struct{}{},
			}
		}
		b := &ReverseNestedAggParam{}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type ReverseNestedAgg func(o ...func(*ReverseNestedAggParam)) map[string]interface{}

func (r ReverseNestedAgg) WithPath(path string) func(*ReverseNestedAggParam) {
	return func(p *ReverseNestedAggParam) {
		if path != "" {
			p.path = path
		}
	}
}

func (r ReverseNestedAgg) WithChildAgg(agg map[string]interface{}) func(*ReverseNestedAggParam) {
	return func(p *ReverseNestedAggParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type ReverseNestedAggParam struct {
	name string
	path string
	aggs map[string]interface{}
}

func (r *ReverseNestedAggParam) Build() map[string]interface{} {
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
