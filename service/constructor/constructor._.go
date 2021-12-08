package constructor

type Constructor struct{
	Compound *Compound
	TermLevel *TermLevel
	FullText *FullText
	Join  *Join
	Common *Common
}

type Compound struct{
	Bool Bool
	Boosting Boosting
	ConstantScore ConstantScore
	DisjunctionMax DisjunctionMax
}

type TermLevel struct{
	Term Term
	Terms Terms
	Range Range
	Ids IDS
}

type FullText struct{
	Match Match
	MatchAll MatchAll
	MatchBoolPrefix MatchBoolPrefix
	MatchPhrase MatchPhrase
	MatchPhrasePrefix MatchPhrasePrefix
	CombinedFields CombinedFields
	MultiMatch MultiMatch
	SimpleQueryString SimpleQueryString
}

type Join struct {
	Nested Nested
}

type Common struct {
	Sort Sort
	Source Source
}

func New() *Constructor{
	return &Constructor{
		Compound: &Compound{
			Bool:newBool(),
			Boosting:newBoosting(),
			ConstantScore:newConstantScore(),
			DisjunctionMax:newDisjunctionMax(),
		},
		TermLevel:&TermLevel{
			Term:newTerm(),
			Terms: newTerms(),
			Range:newRange(),
			Ids: newIDS(),
		},
		FullText:&FullText{
			Match: newMatch(),
			MatchAll:newMatchAll(),
			MatchBoolPrefix:newMatchBoolPrefix(),
			MatchPhrase:newMatchPhrase(),
			MatchPhrasePrefix:newMatchPhrasePrefix(),
			CombinedFields:newCombinedFields(),
			MultiMatch:newMultiMatch(),
			SimpleQueryString:newSimpleQueryString(),
		},
		Join:&Join{
			Nested:newNested(),
		},
		Common: &Common{
			Sort: NewSort(),
			Source: NewSource(),
		},
	}
}