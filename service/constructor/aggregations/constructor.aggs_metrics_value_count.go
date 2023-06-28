package aggregations

func NewValueCount() ValueCount {
	return func(field string) map[string]interface{} {
		return map[string]interface{}{
			"value_count": map[string]interface{}{
				"field": field,
			},
		}
	}
}

type ValueCount func(field string) map[string]interface{}
