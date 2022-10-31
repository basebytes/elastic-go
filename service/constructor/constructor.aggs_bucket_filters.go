package constructor

func newFilters() Filters {
	return func(filter func(*FiltersParam), o ...func(*FiltersParam)) map[string]interface{} {
		b := &FiltersParam{filters: make(map[string]interface{})}
		o = append(o, filter)
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Filters func(filter func(*FiltersParam), o ...func(*FiltersParam)) map[string]interface{}

func (f Filters) WithFilter(name string, query map[string]interface{}) func(*FiltersParam) {
	return func(p *FiltersParam) {
		if name != "" && len(query) > 0 {
			p.filters[name] = query
		}
	}
}

func (f Filters) WithOtherBucket(withOtherBucket bool) func(*FiltersParam) {
	return func(p *FiltersParam) {
		p.otherBucket = withOtherBucket
	}
}

func (f Filters) WithOtherBucketKey(otherBucketKey string) func(*FiltersParam) {
	return func(p *FiltersParam) {
		p.otherBucketKey = otherBucketKey
	}
}

func (f Filters) WithChildAgg(agg map[string]interface{}) func(*FiltersParam) {
	return func(p *FiltersParam) {
		if len(agg) == 0 {
			return
		}
		p.aggs = agg
	}
}

type FiltersParam struct {
	otherBucketKey string
	otherBucket    bool
	aggs           map[string]interface{}
	filters        map[string]interface{}
}

func (f FiltersParam) Build() map[string]interface{} {
	filters := map[string]interface{}{
		"filters": f.filters,
	}
	if f.otherBucket == true {
		filters["other_bucket"] = f.otherBucket
	}
	if f.otherBucketKey != "" {
		filters["other_bucket_key"] = f.otherBucketKey
	}
	res := map[string]interface{}{
		"filters": filters,
	}
	if len(f.aggs) > 0 {
		res["aggs"] = f.aggs
	}
	return res
}
