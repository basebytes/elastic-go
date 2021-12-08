package _examples

import (
	"fmt"
	"testing"
)

func TestCatIndices(t *testing.T){
	result,err:= esClient.Cat.Indices(
		esClient.Cat.Indices.WithIndex(indexAlias+"*"),
		esClient.Cat.Indices.WithColumns("index,health","pri"),
		esClient.Cat.Indices.WithFormat("json"),
	)
	if err==nil{
		fmt.Println(result.ResultString())
	}else{
		fmt.Println(err)
	}
}

func TestPing(t *testing.T){
	result,err:= esClient.Ping()
	if err==nil{
		fmt.Println(result.ResultString())
	}else{
		fmt.Println(err)
	}
}

