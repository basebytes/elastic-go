package _examples

import (
	"fmt"
	"github.com/basebytes/tools"
	"testing"
)

func TestNested(t *testing.T){
	query:=Service.Constructor.FullText.Match(
		"mca.asr.word","可惜")
	source:=Service.Constructor.Common.Source(
		Service.Constructor.Common.Source.WithIncludes("mca.asr.word"),
		)
	join:= Join.Nested("mca.asr",query, Join.Nested.WithSource(source))
	fmt.Println(tools.Encode(join))
}