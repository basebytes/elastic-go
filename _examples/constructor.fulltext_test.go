package _examples

import (
	"fmt"
	"github.com/basebytes/elastic-go/service/constructor/query"
	"github.com/basebytes/tools"
	"testing"
)

func TestMatch(t *testing.T) {
	match := FullText.Match(
		"title", "索引 测试 test",
		FullText.Match.WithAnalyzer("whitespace"),
		FullText.Match.WithOperator("and"),
	)
	fmt.Println(tools.Encode(match))
	match2 := FullText.Match(
		"title", "索引 测试 test",
	)
	fmt.Println(tools.Encode(match2))
}

func TestMatchAll(t *testing.T) {
	match := FullText.MatchAll(
		FullText.MatchAll.WithBoost(2.2),
	)
	fmt.Println(tools.Encode(match))
}

func TestMatchPhrase(t *testing.T) {
	match := FullText.MatchPhrase(
		"title", "索引 测试 test",
		FullText.MatchPhrase.WithAnalyzer("whitespace"),
		FullText.MatchPhrase.WithSlop(2),
	)
	fmt.Println(tools.Encode(match))
}

func TestMultiMatch(t *testing.T) {
	match := FullText.MultiMatch(
		"索引 测试 test",
		FullText.MultiMatch.WithFields("title", "keywords"),
		FullText.MultiMatch.WithType(query.MultiMatchTypeCrossFields),
		FullText.MultiMatch.WithAnalyzer("whitespace"),
		FullText.MultiMatch.WithTieBreaker(0.5),
	)
	fmt.Println(tools.Encode(match))
}

func TestMatchBoolPrefix(t *testing.T) {
	match := FullText.MatchBoolPrefix(
		"title",
		"indexing t",
		FullText.MatchBoolPrefix.WithAnalyzer("whitespace"),
		FullText.MatchBoolPrefix.WithOperator("and"),
	)
	fmt.Println(tools.Encode(match))
}

func TestMatchPhrasePrefix(t *testing.T) {
	match := FullText.MatchPhrasePrefix(
		"title",
		"indexing t",
		FullText.MatchPhrasePrefix.WithAnalyzer("whitespace"),
		FullText.MatchPhrasePrefix.WithZeroTermsQuery("all"),
	)
	fmt.Println(tools.Encode(match))
}

func TestCombinedFields(t *testing.T) {
	match := FullText.CombinedFields(
		[]string{"title"},
		"indexing t",
		FullText.CombinedFields.WithOperator("and"),
		FullText.CombinedFields.WithZeroTermsQuery("all"),
	)
	fmt.Println(tools.Encode(match))
}

func TestSimpleQueryString(t *testing.T) {
	match := FullText.SimpleQueryString(
		"indexing t",
		FullText.SimpleQueryString.WithDefaultOperator("and"),
		FullText.SimpleQueryString.WithFields("title", "keywords"),
		FullText.SimpleQueryString.WithAnalyzer("standard"),
	)
	fmt.Println(tools.Encode(match))
}
