package constructor

import (
	"github.com/basebytes/elastic-go/service/constructor/aggregations"
	"github.com/basebytes/elastic-go/service/constructor/query"
)

type Constructor struct {
	Compound    *Compound
	TermLevel   *TermLevel
	FullText    *FullText
	Join        *Join
	Common      *Common
	Specialized *Specialized
	Aggs        *Aggregations
}

type Compound struct {
	Bool           query.Bool
	Boosting       query.Boosting
	ConstantScore  query.ConstantScore
	DisjunctionMax query.DisjunctionMax
}

type TermLevel struct {
	Term   query.Term
	Terms  query.Terms
	Range  query.Range
	Ids    query.IDS
	Exists query.Exists
}

type FullText struct {
	Match             query.Match
	MatchAll          query.MatchAll
	MatchBoolPrefix   query.MatchBoolPrefix
	MatchPhrase       query.MatchPhrase
	MatchPhrasePrefix query.MatchPhrasePrefix
	CombinedFields    query.CombinedFields
	MultiMatch        query.MultiMatch
	SimpleQueryString query.SimpleQueryString
}

type Join struct {
	Nested query.Nested
}

type Specialized struct {
	Script query.Script
}
type Common struct {
	Sort   query.Sort
	Source query.Source
}

type Aggregations struct {
	Bucket   Bucket
	Metrics  Metrics
	Pipeline Pipeline
}

type Bucket struct {
	DateHistogram aggregations.DateHistogram
	Terms         aggregations.Terms
	Filter        aggregations.Filter
	Filters       aggregations.Filters
	Nested        aggregations.Nested
	ReverseNested aggregations.ReverseNested
	Histogram     aggregations.Histogram
	Range         aggregations.Range
}

type Metrics struct {
	Cardinality aggregations.Cardinality
	Sum         aggregations.Sum
	ValueCount  aggregations.ValueCount
}

type Pipeline struct {
	Script    aggregations.Script
	Selector  aggregations.Selector
	Sort      aggregations.Sort
	SumBucket aggregations.SumBucket
}

func New() *Constructor {
	return &Constructor{
		Compound: &Compound{
			Bool:           query.NewBool(),
			Boosting:       query.NewBoosting(),
			ConstantScore:  query.NewConstantScore(),
			DisjunctionMax: query.NewDisjunctionMax(),
		},
		TermLevel: &TermLevel{
			Term:   query.NewTerm(),
			Terms:  query.NewTerms(),
			Range:  query.NewRange(),
			Ids:    query.NewIDS(),
			Exists: query.NewExists(),
		},
		FullText: &FullText{
			Match:             query.NewMatch(),
			MatchAll:          query.NewMatchAll(),
			MatchBoolPrefix:   query.NewMatchBoolPrefix(),
			MatchPhrase:       query.NewMatchPhrase(),
			MatchPhrasePrefix: query.NewMatchPhrasePrefix(),
			CombinedFields:    query.NewCombinedFields(),
			MultiMatch:        query.NewMultiMatch(),
			SimpleQueryString: query.NewSimpleQueryString(),
		},
		Join: &Join{
			Nested: query.NewNested(),
		},
		Common: &Common{
			Sort:   query.NewSort(),
			Source: query.NewSource(),
		},
		Specialized: &Specialized{
			Script: query.NewScript(),
		},
		Aggs: &Aggregations{
			Bucket: Bucket{
				DateHistogram: aggregations.NewDateHistogram(),
				Terms:         aggregations.NewTerms(),
				Filter:        aggregations.NewFilter(),
				Filters:       aggregations.NewFilters(),
				Nested:        aggregations.NewNested(),
				ReverseNested: aggregations.NewReverseNested(),
				Histogram:     aggregations.NewHistogram(),
				Range:         aggregations.NewRange(),
			},
			Metrics: Metrics{
				Cardinality: aggregations.NewCardinality(),
				Sum:         aggregations.NewSum(),
				ValueCount:  aggregations.NewValueCount(),
			},
			Pipeline: Pipeline{
				Script:    aggregations.NewScript(),
				Selector:  aggregations.NewSelector(),
				Sort:      aggregations.NewSort(),
				SumBucket: aggregations.NewSumBucket(),
			},
		},
	}
}
