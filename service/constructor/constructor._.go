package constructor

type Constructor struct {
	Compound  *Compound
	TermLevel *TermLevel
	FullText  *FullText
	Join      *Join
	Common    *Common
	Aggs      *Aggregations
}

type Compound struct {
	Bool           Bool
	Boosting       Boosting
	ConstantScore  ConstantScore
	DisjunctionMax DisjunctionMax
}

type TermLevel struct {
	Term  Term
	Terms Terms
	Range Range
	Ids   IDS
}

type FullText struct {
	Match             Match
	MatchAll          MatchAll
	MatchBoolPrefix   MatchBoolPrefix
	MatchPhrase       MatchPhrase
	MatchPhrasePrefix MatchPhrasePrefix
	CombinedFields    CombinedFields
	MultiMatch        MultiMatch
	SimpleQueryString SimpleQueryString
}

type Join struct {
	Nested Nested
}

type Common struct {
	Sort   Sort
	Source Source
}

type Aggregations struct {
	Bucket   Bucket
	Metrics  Metrics
	Pipeline Pipeline
}

type Bucket struct {
	DateHistogram    DateHistogram
	TermsAgg         TermsAgg
	Filters          Filters
	NestedAgg        NestedAgg
	ReverseNestedAgg ReverseNestedAgg
	Histogram        Histogram
	RangeAgg         RangeAgg
}

type Metrics struct {
	Sum         Sum
	Cardinality Cardinality
}

type Pipeline struct {
	Script    Script
	SumBucket SumBucket
}

func New() *Constructor {
	return &Constructor{
		Compound: &Compound{
			Bool:           newBool(),
			Boosting:       newBoosting(),
			ConstantScore:  newConstantScore(),
			DisjunctionMax: newDisjunctionMax(),
		},
		TermLevel: &TermLevel{
			Term:  newTerm(),
			Terms: newTerms(),
			Range: newRange(),
			Ids:   newIDS(),
		},
		FullText: &FullText{
			Match:             newMatch(),
			MatchAll:          newMatchAll(),
			MatchBoolPrefix:   newMatchBoolPrefix(),
			MatchPhrase:       newMatchPhrase(),
			MatchPhrasePrefix: newMatchPhrasePrefix(),
			CombinedFields:    newCombinedFields(),
			MultiMatch:        newMultiMatch(),
			SimpleQueryString: newSimpleQueryString(),
		},
		Join: &Join{
			Nested: newNested(),
		},
		Common: &Common{
			Sort:   NewSort(),
			Source: NewSource(),
		},
		Aggs: &Aggregations{
			Bucket: Bucket{
				DateHistogram:    newDateHistogram(),
				TermsAgg:         newTermsAgg(),
				Filters:          newFilters(),
				NestedAgg:        newNestedAgg(),
				ReverseNestedAgg: newReverseNestedAgg(),
				Histogram:        newHistogram(),
				RangeAgg:         newRangeAgg(),
			},
			Metrics: Metrics{
				Sum:         newSum(),
				Cardinality: newCardinality(),
			},
			Pipeline: Pipeline{
				Script:    newScript(),
				SumBucket: newSumBucket(),
			},
		},
	}
}
