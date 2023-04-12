package aggregations

func NewDateHistogram() DateHistogram {
	return func(field string, interval func(*DateHistogramParam), o ...func(*DateHistogramParam)) map[string]interface{} {
		b := &DateHistogramParam{param: map[string]interface{}{
			"field": field,
		}}
		interval(b)
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type DateHistogram func(field string, interval func(*DateHistogramParam), o ...func(*DateHistogramParam)) map[string]interface{}

func (d DateHistogram) WithCalendarInterval(intervals string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if intervals == "" {
			return
		}
		p.param["calendar_interval"] = intervals
		delete(p.param, "fixed_interval")
	}
}

func (d DateHistogram) WithFixedInterval(intervals string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if intervals == "" {
			return
		}
		p.param["fixed_interval"] = intervals
		delete(p.param, "calendar_interval")
	}
}

// 8.4.3无效
func (d DateHistogram) WithFormat(format string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if format == "" {
			return
		}
		p.param["format"] = format
	}
}

func (d DateHistogram) WithTimeZone(timeZone string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if timeZone == "" {
			return
		}
		p.param["time_zone"] = timeZone
	}
}

func (d DateHistogram) WithOffset(offset string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if offset == "" {
			return
		}
		p.param["offset"] = offset
	}
}

func (d DateHistogram) WithMissingValue(value string) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if value == "" {
			return
		}
		p.param["missing"] = value
	}
}

func (d DateHistogram) WithChildAgg(agg map[string]interface{}) func(*DateHistogramParam) {
	return func(p *DateHistogramParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type DateHistogramParam struct {
	param map[string]interface{}
	aggs  map[string]interface{}
}

func (d *DateHistogramParam) Build() map[string]interface{} {

	res := map[string]interface{}{
		"date_histogram": d.param,
	}
	if len(d.aggs) > 0 {
		res["aggs"] = d.aggs
	}
	return res
}
