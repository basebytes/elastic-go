package constructor

func newExists() Exists {
	return func(field string) map[string]interface{} {
		return map[string]interface{}{
			"exists": map[string]interface{}{
				"field": field,
			},
		}
	}
}

type Exists func(field string) map[string]interface{}
