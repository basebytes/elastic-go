package aggregations

func NewCardinality() Cardinality {
	return func(field string, o ...func(*CardinalityParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"cardinality": map[string]interface{}{
					"field": field,
				},
			}
		}
		b := &CardinalityParam{param: map[string]interface{}{"field": field}}
		for _, f := range o {
			f(b)
		}
		return b.Build()
	}
}

type Cardinality func(field string, o ...func(*CardinalityParam)) map[string]interface{}

const maxPrecisionThreshold = 40000

func (c Cardinality) WithPrecisionThreshold(value int) func(*CardinalityParam) {
	return func(p *CardinalityParam) {
		if value < 0 {
			return
		}
		if value > maxPrecisionThreshold {
			value = maxPrecisionThreshold
		}
		p.param["precision_threshold"] = value
	}
}

func (c Cardinality) WithMissingValue(value interface{}) func(*CardinalityParam) {
	return func(p *CardinalityParam) {
		if value == nil {
			return
		}
		p.param["missing"] = value
	}
}

type CardinalityParam struct {
	param map[string]interface{}
}

func (c CardinalityParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"cardinality": c.param,
	}
}
