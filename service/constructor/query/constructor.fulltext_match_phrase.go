package query

import "strings"

func NewMatchPhrase() MatchPhrase {
	return func(field, text string, o ...func(*MatchPhraseParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"match_phrase": map[string]interface{}{
					field: text,
				},
			}
		}
		m := &MatchPhraseParam{field: field, param: map[string]interface{}{
			"query": text,
		}}
		for _, f := range o {
			f(m)
		}
		return m.Build()
	}
}

type MatchPhrase func(field, text string, o ...func(*MatchPhraseParam)) map[string]interface{}

//default  analysis.analyzer.default or standard analyzer
func (m MatchPhrase) WithAnalyzer(analyzer string) func(*MatchPhraseParam) {
	return func(p *MatchPhraseParam) {
		if analyzer != "" {
			p.param["analyzer"] = analyzer
		}
	}
}

//default none
func (m MatchPhrase) WithZeroTermsQuery(zeroTermsQuery string) func(*MatchPhraseParam) {
	return func(p *MatchPhraseParam) {
		switch strings.ToLower(zeroTermsQuery) {
		case "all":
			p.param["zero_terms_query"] = "all"
		}
	}
}

//default 1.0
func (m MatchPhrase) WithBoost(boost float32) func(*MatchPhraseParam) {
	return func(p *MatchPhraseParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

//default 0
func (m MatchPhrase) WithSlop(slop int) func(*MatchPhraseParam) {
	return func(p *MatchPhraseParam) {
		if slop > 0 {
			p.param["slop"] = slop
		}
	}
}

func (m MatchPhrase) WithName(name string) func(*MatchPhraseParam) {
	return func(p *MatchPhraseParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

type MatchPhraseParam struct {
	field string
	param map[string]interface{}
}

//{
//	"match_phrase":{
//		<field>:{
//			"query":<text>,
//			...
//		}
//	}
//}
//
//or
//
//{
//	"match_phrase":{
//		<field>:<text>,
//	}
//}
func (m *MatchPhraseParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"match_phrase": map[string]interface{}{
			m.field: m.param,
		},
	}
}
