package entity

type EsQueryResult struct {
	EsError	`mapstructure:",squash"`
	Took     int32
	TimedOut bool `mapstructure:"timed_out"`
	Shards   *struct {
		Total      uint32
		Successful uint32
		Skipped    uint32
		Failed     uint32
	} `mapstructure:"_shards"`
	Hits *Hits
	Aggs Aggregations `mapstructure:"aggregations"`
}

//aggs
type Aggregations map[string]*AggregationsResult

type AggregationsResult struct {
	Buckets *[]*BucketItem
	Other   map[string]interface{} `mapstructure:",remain"`
}

type BucketItem struct {
	Key      string
	DocCount int64                  `mapstructure:"doc_count"`
	Other    map[string]interface{} `mapstructure:",remain"`
}
// MetricsAggTopHits
type MetricsAggTopHitsResultItem struct {
	Hits *Hits
}

type EsAggregationRangeResult struct {
	From       uint64
	To         uint64
	DocCount   uint32 `mapstructure:"doc_count"`
	TopDocHits struct {
		Hits *Hits
	} `mapstructure:"top_doc_hits,omitempty"`
}

