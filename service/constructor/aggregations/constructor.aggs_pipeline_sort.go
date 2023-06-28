package aggregations

func NewSort() Sort {
	return func(sort func(*SortParam), o ...func(*SortParam)) map[string]interface{} {
		b := &SortParam{}
		sort(b)
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Sort func(sort func(*SortParam), o ...func(*SortParam)) map[string]interface{}

func (s Sort) WithAsc(name string) func(param *SortParam) {
	return func(param *SortParam) {
		param.aggOrder(name, orderAsc)
	}
}

func (s Sort) WithDesc(name string) func(param *SortParam) {
	return func(param *SortParam) {
		param.aggOrder(name, orderDesc)
	}
}

func (s Sort) WithFrom(from int) func(param *SortParam) {
	return func(param *SortParam) {
		if from > 0 {
			param.from = from
		}
	}
}

func (s Sort) WithSize(size int) func(param *SortParam) {
	return func(param *SortParam) {
		if size > 0 {
			param.size = size
		}
	}
}

type SortParam struct {
	sort []interface{}
	from int
	size int
}

func (s SortParam) Build() map[string]interface{} {
	sort := map[string]interface{}{
		"sort": s.sort,
	}
	if s.from > 0 {
		sort["from"] = s.from
	}
	if s.size > 0 {
		sort["size"] = s.size
	}
	return map[string]interface{}{
		"bucket_sort": sort,
	}
}

func (s SortParam) aggOrder(name, order string) {
	if name != "" {
		s.sort = append(s.sort, map[string]interface{}{
			name: map[string]interface{}{
				"order": order,
			},
		})
	}
}

const (
	orderAsc  = "asc"
	orderDesc = "desc"
)
