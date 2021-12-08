package _examples

import (
	"github.com/basebytes/elastic-go/client"
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/elastic-go/service"
	"github.com/basebytes/elastic-go/service/constructor"
	"github.com/basebytes/tools"
	"time"
)

const (
	dateFormat = "20060102"
	indexAlias = "test"
	testTableJson=`{
  "settings": {
    "index": {
      "number_of_shards": 1,
      "number_of_replicas": 0,
      "max_inner_result_window":500,
      "max_result_window":500,
      "translog":{
        "durability":"async",
        "sync_interval":"150s",
        "flush_threshold_size":"1024mb"
      },
      "merge": {
        "scheduler": {
          "max_thread_count": "1"
        }
      }
    }
  ,"refresh_interval": "1s"
  },
  "mappings": {
    "dynamic": false,
    "properties" : {
      "id" : {
        "type" : "keyword"
      },
      "title" : {
        "type" : "text"
      },
      "published" : {
        "type" : "integer"
      },
      "duration" : {
        "type" : "integer"
      },
      "keywords" : {
        "type" : "keyword"
      }
    }
  }
}`
)

var (
	dateStr=time.Now().Format(dateFormat)
	currentIndexName= indexAlias + dateStr

	err error
	esClient *client.Client
	Document *api.Document
	Index *api.Indexer
	Searcher *api.Searcher

	Service *service.Service

	TermLevel *constructor.TermLevel
	Compound *constructor.Compound
	FullText *constructor.FullText
	Join *constructor.Join
)

func init(){
	esClient, err =client.NewDefaultClient()
	if err !=nil{
		panic(err)
	}
	Document = esClient.Document
	Index = esClient.Indexer
	Searcher = esClient.Searcher
	Service,err= service.NewDefaultService()
	if err!=nil{
		panic(err)
	}
	TermLevel =Service.Constructor.TermLevel
	Compound =Service.Constructor.Compound
	FullText =Service.Constructor.FullText
	Join =Service.Constructor.Join
}

type TestDoc struct{
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Published int64 `json:"published,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	action string `json:"-"`
	index string `json:"-"`
}

func (t *TestDoc)BulkAction() *api.BulkAction{
	action:=&api.BulkAction{}
	actionItem:=&api.BulkItem{
		Index: t.index,
		Id:    t.Id,
	}
	switch t.action {
	case "create":action.Create=actionItem
	case "delete":action.Delete=actionItem
	case "index":action.Index=actionItem
	case "update":action.Update=actionItem
	default:return nil
	}
	return action
}

func (t *TestDoc)Content() []byte{
	return tools.EncodeBytes(t)
}

func (t *TestDoc)GetUUID() string{
	return t.Id
}