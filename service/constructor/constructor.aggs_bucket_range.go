package constructor

func newRangeAgg() RangeAgg {
	return func(field string, dataRange func(*RangeAggParam), o ...func(*RangeAggParam)) map[string]interface{} {
		b := &RangeAggParam{param: map[string]interface{}{
			"field": field,
		}}
		dataRange(b)
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type RangeAgg func(field string, dataRange func(*RangeAggParam), o ...func(*RangeAggParam)) map[string]interface{}

func (r RangeAgg) WithKeyed() func(param *RangeAggParam) {
	return func(p *RangeAggParam) {
		p.param["keyed"] = true
	}
}

func (r RangeAgg) WithRange(key string, from, to interface{}) func(param *RangeAggParam) {
	return func(p *RangeAggParam) {
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

func (r RangeAgg) WithChildAgg(agg map[string]interface{}) func(*RangeAggParam) {
	return func(p *RangeAggParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type RangeAggParam struct {
	ranges []map[string]interface{}
	param  map[string]interface{}
	aggs   map[string]interface{}
}

func (r RangeAggParam) Build() map[string]interface{} {
	r.param["ranges"] = r.ranges
	res := map[string]interface{}{
		"range": r.param,
	}
	if len(r.aggs) > 0 {
		res["aggs"] = r.aggs
	}
	return res
}
