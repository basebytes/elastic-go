package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/service/constructor/query"
	"github.com/basebytes/tools"
	"testing"
)

func TestTerm(t *testing.T) {
	term := TermLevel.Term("keywords", "test",
		TermLevel.Term.WithBoost(1.2),
	)
	fmt.Println(tools.Encode(term))
	term2 := TermLevel.Term("keywords", "test1")//TermLevel.Term.WithBoost(1.2),

	fmt.Println(tools.Encode(term2))
}

func TestTerms(t *testing.T) {
	terms := TermLevel.Terms(
		"keywords",
		TermLevel.Terms.WithValue("test", "测试"),
		TermLevel.Terms.WithBoost(1.5),
	)
	fmt.Println(tools.Encode(terms))
	terms2 := TermLevel.Terms(
		"keywords",
		TermLevel.Terms.WithLookUp(indexAlias,
			"ad11bd7e-f5fe-48e7-990b-563294ac6499",
			"keywords",
			"",
		),
		TermLevel.Terms.WithName("v"),
		//TermLevel.Terms.WithBoost(1.5),
	)
	fmt.Println(tools.Encode(terms2))
}

func TestRange(t *testing.T) {
	term := TermLevel.Range("keywords",
		TermLevel.Range.WithCompareOperate(query.CompareOperatorGT, "t"),
		TermLevel.Range.WithBoost(1.2),
	)
	fmt.Println(tools.Encode(term))
}

func TestIds(t *testing.T) {
	term := TermLevel.Ids([]string{"1", "2", "3"},
		//TermLevel.Range.WithCompareOperate(constructor.CompareOperatorGT,"t"),
		TermLevel.Ids.WithBoost(1.2),
		TermLevel.Ids.WithName("aa"),
	)
	fmt.Println(tools.Encode(term))
}
