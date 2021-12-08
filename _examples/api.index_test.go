package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/basebytes/tools"
	"testing"
)


func TestIndexExists(t *testing.T){
	result,err:= Index.Exists(
		Index.Exists.WithIndex(currentIndexName),
	)
	if err !=nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println(result.StatusCode)
}

func TestIndexCreate(t *testing.T){
	schema:=[]byte(testTableJson)
	resp,e:= Index.Create(currentIndexName, Index.Create.WithBody(&schema))
	if e!=nil{
		panic(e)
	}
	var operateResult = new(entity.OperateResult)
	e=resp.Result(operateResult)
	if e!=nil{
		panic(e)
	}
	fmt.Println(operateResult.Acknowledged)
}

func TestIndexAliasExists(t *testing.T){
	resp,e:= Index.AliasGet(
		Index.AliasGet.WithAlias(indexAlias),
	)
	if e!=nil{
		panic(e)
	}
	var indexAlias=new(entity.IndexAlias)
	e=resp.Result(indexAlias)
	if e!=nil{
		panic(e)
	}
	fmt.Println(tools.Encode(indexAlias))
}

func TestIndexAliases(t *testing.T){
	resp,e:= Index.Aliases(
		[]*api.AliasAction{
			{Add: &api.ActionParam{Index: currentIndexName,Alias: indexAlias}},
		},
	)
	if e!=nil{
		panic(e)
	}
	var operateResult = new(entity.OperateResult)
	e=resp.Result(operateResult)
	if e!=nil{
		panic(e)
	}
	fmt.Println(tools.Encode(operateResult))
}

func TestIndexDelete(t *testing.T){
	resp,e:= Index.Delete([]string{currentIndexName})
	if e!=nil{
		panic(e)
	}
	var operateResult = new(entity.OperateResult)
	e=resp.Result(operateResult)
	if e!=nil{
		panic(e)
	}
	fmt.Println(operateResult.Acknowledged)
}

func TestIndexGetMapping(t *testing.T){
	resp,e:= Index.GetMapping(
		//Indexer.GetMapping.WithIndex(currentIndexName),
	)
	if e!=nil{
		panic(e)
	}
	var indexMapping =new(entity.IndexMapping)
	e=resp.Result(indexMapping)
	if e!=nil{
		panic(e)
	}
	fmt.Println(tools.Encode(indexMapping))
}

func TestIndexSettingsGet(t *testing.T){
	resp,e:= Index.SettingsGet(
		Index.SettingsGet.WithIndex(currentIndexName),
	)
	if e!=nil{
		panic(e)
	}
	var indexSetting =new(entity.IndexSetting)
	e=resp.Result(indexSetting)
	if e!=nil{
		panic(e)
	}
	fmt.Println(tools.Encode(indexSetting))
}
