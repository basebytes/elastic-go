package constructor

func newSumBucket() SumBucket {
	return func(path string) map[string]interface{} {
		return map[string]interface{}{
			"sum_bucket": map[string]interface{}{
				"buckets_path": path,
			},
		}
	}
}

type SumBucket func(path string) map[string]interface{}
