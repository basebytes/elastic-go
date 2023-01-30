package constructor

func newTermsAgg() TermsAgg {
	return func(field string, o ...func(*TermsAggParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"terms": map[string]interface{}{
					"field": field,
				},
			}
		}
		b := &TermsAggParam{param: map[string]interface{}{"field": field}}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type TermsAgg func(field string, o ...func(*TermsAggParam)) map[string]interface{}

func (t TermsAgg) WithSize(size int) func(param *TermsAggParam) {
	return func(p *TermsAggParam) {
		if size <= 0 {
			return
		}
		p.param["size"] = size
	}
}

func (t TermsAgg) WithChildAgg(agg map[string]interface{}) func(*TermsAggParam) {
	return func(p *TermsAggParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

func (t TermsAgg) WithMissingValue(value interface{}) func(*TermsAggParam) {
	return func(p *TermsAggParam) {
		if value == nil {
			return
		}
		p.param["missing"] = value
	}
}

type TermsAggParam struct {
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (t TermsAggParam) Build() map[string]interface{} {
	res := map[string]interface{}{
		"terms": t.param,
	}
	if len(t.aggs) > 0 {
		res["aggs"] = t.aggs
	}
	return res
}
