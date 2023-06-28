package aggregations

func NewFilter() Filter {
	return func(filterQuery map[string]interface{}, o ...func(*FilterParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"filter": filterQuery,
			}
		}
		b := &FilterParam{query: filterQuery}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Filter func(filterQuery map[string]interface{}, o ...func(*FilterParam)) map[string]interface{}

func (f Filter) WithChildAgg(agg map[string]interface{}) func(*FilterParam) {
	return func(p *FilterParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type FilterParam struct {
	query map[string]interface{}
	aggs  map[string]interface{}
}

func (f FilterParam) Build() map[string]interface{} {
	res := map[string]interface{}{
		"filter": f.query,
	}
	if len(f.aggs) > 0 {
		res["aggs"] = f.aggs
	}
	return res
}
