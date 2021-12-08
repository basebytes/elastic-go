package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/tools"
	"testing"
)



func TestBulk(t *testing.T){
	resp,err:= Document.Bulk(
		[]api.Doc{
			&TestDoc{action: "index",Id:tools.RandUUID(),Title:"索引测试",Keywords:[]string{"索引","测试"}},
		},
		Document.Bulk.WithIndex(currentIndexName),
	)
	if err!=nil{
		panic(err)
	}
	fmt.Println(resp.ResultString())
}

func TestGet(t *testing.T){
	resp,err:= Document.Get(
		indexAlias,
		"f3e59b33-70d7-4aa3-b0bb-163995ae7a27",
	)
	if err!=nil{
		panic(err)
	}
	var result TestDoc
	fmt.Println(resp.Result(&result),tools.Encode(result))
}

func TestIndex(t *testing.T){
	resp,err:= Document.Index(
		currentIndexName,
		&TestDoc{Id: tools.RandUUID(),Title:"索引测试1",Keywords:[]string{"索引1","测试1"}},
	)
	if err!=nil{
		panic(err)
	}
	fmt.Println(resp.ResultString())
}

