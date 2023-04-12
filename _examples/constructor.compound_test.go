package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/service/constructor/query"
	"github.com/basebytes/tools"
	"testing"
)

func TestBool(t *testing.T) {
	term := TermLevel.Term("keywords", "test",
		TermLevel.Term.WithBoost(1.2),
		TermLevel.Term.WithName("v"),
	)
	term2 := TermLevel.Term("keywords", "test1",
		TermLevel.Term.WithName("v2"),
	)
	boolQuery := Compound.Bool(
		Compound.Bool.WithClause(query.ClauseTypeShould, term, term2),
		//Compound.Bool.WithClause(constructor.ClauseTypeShould,term2),
	)
	fmt.Println(tools.Encode(boolQuery))
}

func TestConstantScore(t *testing.T) {
	term := TermLevel.Term("keywords", "test",
		TermLevel.Term.WithBoost(1.2),
		TermLevel.Term.WithName("v"),
	)
	constantScoreQuery := Compound.ConstantScore(
		term,
		//Compound.ConstantScore.WithBoost(2.2),
		Compound.ConstantScore.WithName("cs"),
	)
	fmt.Println(tools.Encode(constantScoreQuery))
}

func TestBoosting(t *testing.T) {
	term := TermLevel.Term("keywords", "test",
		TermLevel.Term.WithBoost(1.2),
		TermLevel.Term.WithName("v"),
	)
	term2 := TermLevel.Term("keywords", "test1",
		TermLevel.Term.WithName("v2"),
	)
	boostingQuery := Compound.Boosting(
		term,
		term2,
		0.2,
		//Compound.Boosting.WithBoost(2.2),
		Compound.Boosting.WithName("b"),
	)
	fmt.Println(tools.Encode(boostingQuery))
}

func TestDisjunctionMax(t *testing.T) {
	term := TermLevel.Term("keywords", "test",
		TermLevel.Term.WithBoost(1.2),
		TermLevel.Term.WithName("v"),
	)
	term2 := TermLevel.Term("keywords", "test1",
		TermLevel.Term.WithName("v2"),
	)
	disjunctionMaxQuery := Compound.DisjunctionMax(
		[]map[string]interface{}{term,
			term2,
		},
		Compound.DisjunctionMax.WithBoost(2.2),
		Compound.DisjunctionMax.WithName("d"),
	)
	fmt.Println(tools.Encode(disjunctionMaxQuery))
}
