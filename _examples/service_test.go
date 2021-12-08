package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/elastic-go/service/helper"
	"github.com/basebytes/tools"
	"testing"
)

func  TestServiceBulk(t *testing.T){
	res,err:=Service.Bulk(indexAlias,[]api.Doc{
		&TestDoc{action:"index",Id:tools.RandUUID(),Title:"批量索引测试",Keywords:[]string{"批量","索引","测试"}},
		&TestDoc{action:"index",Id:tools.RandUUID(),Title:"批量索引测试1",Keywords:[]string{"批量1","索引1","测试1"}},
	})
	if err==nil{
		fmt.Println(tools.Encode(res))
	}else{
		fmt.Println(err)
	}
}

func TestIndexMapping(t *testing.T){
	res,err:=Service.GetIndexMapping(indexAlias +"*")
	if err==nil{
		//fmt.Println(tools.Encode(res))
		fmt.Println(tools.Encode(res.GetFields()))
	}else{
		fmt.Println(err)
	}
}

func TestIndexSetting(t *testing.T){
	res,err:=Service.GetIndexSetting(indexAlias,
		//"index.merge.scheduler.max_thread_count",
		//"index.max_inner_result_window",
		//"index.translog.flush_threshold_size",
		)
	if err==nil{
		fmt.Println(tools.Encode(res))
		//fmt.Println(tools.Encode(res.GetSettings("index.translog")))
	}else{
		fmt.Println(err)
	}
}

func TestServiceIndexExists(t *testing.T){
	exists,err:=Service.IndexExists(currentIndexName)
	fmt.Println(exists,err)
}

func TestGetIndexAlias(t *testing.T){
	res,err:=Service.GetIndexAlias(
		//nil,
		[]string{indexAlias},
		nil,
		)
	if err==nil{
		fmt.Println(tools.Encode(res))
		fmt.Println(res.GetIndexByAliasName(indexAlias))
	}else{
		fmt.Println("err:",err)
	}
}

func TestCreateIndex(t *testing.T){
	tb:=[]byte(testTableJson)
	res,err:=Service.CreateIndex(currentIndexName,&tb)
	if err==nil{
		fmt.Println(tools.Encode(res))
	}else{
		fmt.Println("err:",err)
	}
}

func TestDeleteIndex(t *testing.T) {
	res,err:=Service.DeleteIndex(currentIndexName)
	if err==nil{
		fmt.Println(tools.Encode(res))
	}else{
		fmt.Println("err:",err)
	}
}

func TestGetIndexNames(t *testing.T) {
	res,err:=Service.GetIndexNames(indexAlias+"*")
	if err==nil{
		fmt.Println(res)
	}else{
		fmt.Println("err:",err)
	}
}

func TestUpdateIndexAlias(t *testing.T) {
	add:=helper.NewAddAction([]string{currentIndexName},[]string{indexAlias},false)
	remove:=helper.NewRemoveAction([]string{currentIndexName+"_01"},[]string{indexAlias},false)
	res,err:=Service.UpdateIndexAlias(add,remove)
	if err==nil{
		fmt.Println(tools.Encode(res))
	}else{
		fmt.Println("err:",err)
	}
}

func TestServicePing(t *testing.T){
	s:=Service.Ping()
	fmt.Println(s)
}