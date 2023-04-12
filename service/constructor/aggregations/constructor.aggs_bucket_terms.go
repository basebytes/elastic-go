package aggregations

func NewTerms() Terms {
	return func(field string, o ...func(*TermsParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"terms": map[string]interface{}{
					"field": field,
				},
			}
		}
		b := &TermsParam{param: map[string]interface{}{"field": field}}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Terms func(field string, o ...func(*TermsParam)) map[string]interface{}

func (t Terms) WithSize(size int) func(param *TermsParam) {
	return func(p *TermsParam) {
		if size <= 0 {
			return
		}
		p.param["size"] = size
	}
}

func (t Terms) WithChildAgg(agg map[string]interface{}) func(*TermsParam) {
	return func(p *TermsParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

func (t Terms) WithMissingValue(value interface{}) func(*TermsParam) {
	return func(p *TermsParam) {
		if value == nil {
			return
		}
		p.param["missing"] = value
	}
}

type TermsParam struct {
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (t TermsParam) Build() map[string]interface{} {
	res := map[string]interface{}{
		"terms": t.param,
	}
	if len(t.aggs) > 0 {
		res["aggs"] = t.aggs
	}
	return res
}
