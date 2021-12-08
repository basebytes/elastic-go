package api

import "net/http"

type API struct {
	Document *Document
	Searcher *Searcher
	Indexer  *Indexer
	Cat      *Cat
	Ping     Ping
}

type Document struct {
	Index DocumentIndex
	Bulk  DocumentBulk
	Get	Get
}

type Searcher struct {
	Search Search
}

type Indexer struct {
	Create      IndexCreate
	Delete      IndexDelete
	GetMapping  IndexGetMapping
	Exists      IndexExists
	AliasExists AliasExists
	AliasGet    AliasGet
	Aliases Aliases
	SettingsGet SettingsGet
}

type Cat struct {
	Indices CatIndices
}

func New(t Transport) *API {
	return &API{
		Document: &Document{
			Index: newDocumentIndexFunc(t),
			Bulk:  newDocumentBulkFunc(t),
			Get:newDocumentGetFunc(t),
		},
		Searcher: &Searcher{
			Search:newSearchFunc(t),
		},
		Indexer: &Indexer{
			Create:      newIndexCreteFunc(t),
			Delete:      newIndexDelete(t),
			GetMapping:  newIndexGetMappingFunc(t),
			Exists:      newIndexExistsFunc(t),
			AliasExists: newAliasExistsFunc(t),
			AliasGet:    newAliasGetFunc(t),
			SettingsGet: newSettingsGet(t),
			Aliases: newAliasesFunc(t),
		},
		Cat: &Cat{
			Indices: newCatIndices(t),
		},
		Ping :newPingFunc(t),
	}
}

var (
	MethodPostFunc=func() string{return http.MethodPost}
	MethodPutFunc=func() string{return http.MethodPut}
	MethodGetFunc=func() string{return http.MethodGet}
	MethodDeleteFunc=func() string{return http.MethodDelete}
	MethodHeadFunc=func() string{return http.MethodHead}

)