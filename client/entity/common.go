package entity

type EsError struct {
	Error string `json:"error,omitempty"`
	Status int32 `json:"status,omitempty"`
}

type ErrorDetail struct{
	RootCause *[]*struct {
		ErrType string `json:"type" mapstructure:"type"`
		Reason  string `json:"reason"`
	} `json:"root_cause" mapstructure:"root_cause"`
	ErrType string `json:"type" mapstructure:"type"`
	Reason  string `json:"reason"`
}

type Hits struct {
	Total *struct {
		Value    int64	`json:"value"`
		Relation string	`json:"relation"`
	}	`json:"total"`
	MaxScore float32 `json:"max_score" mapstructure:"max_score"`
	Hits     *[]*Hit	`json:"hits,omitempty" mapstructure:"hits"`
}

type Hit struct {
	Meta   `mapstructure:",squash"`
	Source map[string]interface{} `json:"_source,omitempty" mapstructure:"_source"`
	Sort   *[]interface{}         `json:"sort,omitempty"`
	InnerHits map[string]*InnerHitItem	`json:"inner_hits,omitempty" mapstructure:"inner_hits"`
}

type Meta struct {
	Index string  `json:"index,omitempty" mapstructure:"_index"`
	Type  string  `json:"type,omitempty" mapstructure:"_type"`
	Id    string  `json:"id,omitempty" mapstructure:"_id"`
	Score float32 `json:"score,omitempty" mapstructure:"_score"`
	Version uint32 `json:"version,omitempty" mapstructure:"_version"`
	Nested	*Nested	`json:"_nested,omitempty" mapstructure:"_nested"`
}

type Nested struct {
	Field string	`json:"field,omitempty"`
	Offset int 		`json:"offset,omitempty"`
}

type InnerHitItem struct {
	Hits *Hits `json:"hits,omitempty"`
}

type OperateResult struct {
	EsError	`mapstructure:",squash"`
	Acknowledged       bool   `json:"acknowledged,omitempty"`
	ShardsAcknowledged bool   `json:"shards_acknowledged,omitempty" mapstructure:"shards_acknowledged"`
	Index              string `json:"index,omitempty"`
}

type Schema struct {
	MappingItem
	SettingItem
}