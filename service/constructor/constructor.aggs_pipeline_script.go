package constructor

func newScript() Script {
	return func(name string, path map[string]string, script string) map[string]interface{} {
		return map[string]interface{}{
			name: map[string]interface{}{
				"bucket_script": map[string]interface{}{
					"buckets_path": path,
					"script":       script,
				},
			},
		}
	}
}

type Script func(name string, path map[string]string, script string) map[string]interface{}
