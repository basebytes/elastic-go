package entity

type DocumentBulkResult struct {
	Took   int32	`json:"took"`
	Errors bool		`json:"errors"`
	Items  *[]map[string]*DocumentBulkResultItem `json:"items"`
}

type DocumentBulkResultItem struct {
	DocumentIndexResult	`mapstructure:",squash"`
	Status uint32	`json:"status"`
	Error  *struct {
		ErrType   string `json:"status" mapstructure:"type"`
		Reason    string	`json:"reason"`
		IndexUUid string `json:"uuid" mapstructure:"index_uuid"`
		Shard     string	`json:"shard"`
		Index     string	`json:"index"`
	} `json:"error,omitempty"`
	//ignore root_cause right now
}

type DocumentIndexResult struct {
	Meta   `mapstructure:",squash"`
	Result  string	`json:"result"`
	Shards  *struct {
		Total      uint32	`json:"total"`
		Successful uint32	`json:"successful"`
		Failed     uint32	`json:"failed"`
	} `json:"result" mapstructure:"_shards"`
	SeqNo       uint32 `json:"seq_no" mapstructure:"_seq_no"`
	PrimaryTerm uint32 `json:"primary_term" mapstructure:"_primary_term"`
}