package entity


//cat indices
type IndexInfo struct {
	Health       string	 `json:"health,omitempty"`
	Status       string	 `json:"status,omitempty"`
	Index        string	 `json:"index,omitempty"`
	UUid         string	 `json:"uuid,omitempty"`
	Pri          string	 `json:"pri,omitempty"`
	Rep          string	 `json:"rep,omitempty"`
	DocsCount    string `json:"docs.count,omitempty"`
	DocsDeleted  string `json:"docs.deleted,omitempty"`
	StoreSize    string `json:"store.size,omitempty"`
	PriStoreSize string `json:"pri.store.size,omitempty"`
}