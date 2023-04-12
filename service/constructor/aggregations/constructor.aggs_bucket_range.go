package aggregations

func NewRange() Range {
	return func(field string, dataRange func(*RangeParam), o ...func(*RangeParam)) map[string]interface{} {
		b := &RangeParam{param: map[string]interface{}{
			"field": field,
		}}
		dataRange(b)
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Range func(field string, dataRange func(*RangeParam), o ...func(*RangeParam)) map[string]interface{}

func (r Range) WithKeyed() func(param *RangeParam) {
	return func(p *RangeParam) {
		p.param["keyed"] = true
	}
}

func (r Range) WithRange(key string, from, to interface{}) func(param *RangeParam) {
	return func(p *RangeParam) {
		dataRange := make(map[string]interface{})
		if key != "" {
			dataRange["key"] = key
		}
		if from != nil {
			dataRange["from"] = from
		}
		if to != nil {
			dataRange["to"] = to
		}
		p.ranges = append(p.ranges, dataRange)
	}
}

func (r Range) WithChildAgg(agg map[string]interface{}) func(*RangeParam) {
	return func(p *RangeParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type RangeParam struct {
	ranges []map[string]interface{}
	param  map[string]interface{}
	aggs   map[string]interface{}
}

func (r RangeParam) Build() map[string]interface{} {
	r.param["ranges"] = r.ranges
	res := map[string]interface{}{
		"range": r.param,
	}
	if len(r.aggs) > 0 {
		res["aggs"] = r.aggs
	}
	return res
}
