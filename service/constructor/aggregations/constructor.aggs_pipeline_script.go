package aggregations

func NewScript() Script {
	return func(path map[string]string, script interface{}) map[string]interface{} {
		return map[string]interface{}{
			"bucket_script": map[string]interface{}{
				"buckets_path": path,
				"script":       script,
			},
		}
	}
}

type Script func(path map[string]string, script interface{}) map[string]interface{}
