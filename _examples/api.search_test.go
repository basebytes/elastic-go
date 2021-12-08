package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/basebytes/tools"
	"testing"
)

func TestSearch(t *testing.T){
	body:=[]byte(`{
  "query": {
    "match": {
      "keywords": "test"
    }
  },
  "sort": [
    {
      "id": {
        "order": "asc"
      }
    }
  ]
}`)
	resp,err:= Searcher.Search(
		Searcher.Search.WithIndex(indexAlias),
		Searcher.Search.WithBody(&body),
		//Searcher.Search.WithTrackTotalHits(true),
	)
	if err!=nil{
		panic(err)
	}
	var result entity.EsQueryResult
	resp.Result(&result)
	fmt.Println(tools.Encode(result))
	fmt.Println(resp.ResultString())
}