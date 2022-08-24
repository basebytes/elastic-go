package api

type Doc interface {
	BulkAction() *BulkAction
	Content() []byte
	GetUUID() string
}

type BulkAction struct {
	Create *BulkItem `json:"create,omitempty"`
	Delete *BulkItem `json:"delete,omitempty"`
	Index  *BulkItem `json:"index,omitempty"`
	Update *BulkItem `json:"update,omitempty"`
}

type BulkItem struct {
	Index           string `json:"_index,omitempty"`
	Id              string `json:"_id,omitempty"`
	RetryOnConflict int    `json:"retry_on_conflict,omitempty"`
}
