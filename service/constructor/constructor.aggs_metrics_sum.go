package constructor

func newSum() Sum {
	return func(field string, o ...func(*SumParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"sum": map[string]interface{}{
					"field": field,
				},
			}
		}
		b := &SumParam{param: map[string]interface{}{"field": field}}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Sum func(field string, o ...func(*SumParam)) map[string]interface{}

// 不做合法性校验
func (s Sum) WithMissingValue(value interface{}) func(*SumParam) {
	return func(p *SumParam) {
		p.param["missing"] = value
	}
}

type SumParam struct {
	param map[string]interface{}
}

func (s SumParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"sum": s.param,
	}
}
