package aggregations

func NewHistogram() Histogram {
	return func(field string, interval interface{}, o ...func(param *HistogramParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"histogram": map[string]interface{}{
					"field":    field,
					"interval": interval,
				},
			}
		}
		b := &HistogramParam{param: map[string]interface{}{
			"field":    field,
			"interval": interval,
		}}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Histogram func(field string, interval interface{}, o ...func(param *HistogramParam)) map[string]interface{}

func (h Histogram) WithMinDocCount(count int) func(param *HistogramParam) {
	return func(p *HistogramParam) {
		if count > 0 {
			p.param["min_doc_count"] = count
		}
	}
}

func (h Histogram) WithOffset(offset interface{}) func(param *HistogramParam) {
	return func(p *HistogramParam) {
		if offset != nil {
			p.param["offset"] = offset
		}
	}
}

func (h Histogram) WithMissingValue(value interface{}) func(*HistogramParam) {
	return func(p *HistogramParam) {
		if value != nil {
			return
		}
		p.param["missing"] = value
	}
}

func (h Histogram) WithChildAgg(agg map[string]interface{}) func(*HistogramParam) {
	return func(p *HistogramParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type HistogramParam struct {
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (h *HistogramParam) Build() map[string]interface{} {
	res := map[string]interface{}{
		"histogram": h.param,
	}
	if len(h.aggs) > 0 {
		res["aggs"] = h.aggs
	}
	return res
}
