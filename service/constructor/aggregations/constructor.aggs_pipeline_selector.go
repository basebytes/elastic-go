package aggregations

func NewSelector() Selector {
	return func(path map[string]string, script interface{}) map[string]interface{} {
		return map[string]interface{}{
			"bucket_selector": map[string]interface{}{
				"buckets_path": path,
				"script":       script,
			},
		}
	}
}

type Selector func(path map[string]string, script interface{}) map[string]interface{}
