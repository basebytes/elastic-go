package constructor

import "strings"

func newMatchPhrasePrefix() MatchPhrasePrefix{
	return func(field, text string, o ...func(*MatchPhrasePrefixParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"match_phrase_prefix":map[string]interface{}{
					field:text,
				},
			}
		}
		m:=&MatchPhrasePrefixParam{field:field,param: map[string]interface{}{
			"query":text,
		}}
		for _,f:=range o{
			f(m)
		}
		return m.Build()
	}
}

type MatchPhrasePrefix func(field,text string,o ...func(*MatchPhrasePrefixParam)) map[string]interface{}


//default  analysis.analyzer.default or standard analyzer
func (m MatchPhrasePrefix) WithAnalyzer(analyzer string)func(*MatchPhrasePrefixParam){
	return func(p *MatchPhrasePrefixParam) {
		if analyzer!=""{
			p.param["analyzer"]=analyzer
		}
	}
}

//default 1.0
func (m MatchPhrasePrefix) WithBoost(boost float32) func (*MatchPhrasePrefixParam){
	return func(p *MatchPhrasePrefixParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}
//default none
func (m MatchPhrasePrefix) WithZeroTermsQuery(zeroTermsQuery string)func(*MatchPhrasePrefixParam){
	return func(p *MatchPhrasePrefixParam) {
		switch strings.ToLower(zeroTermsQuery) {
		case "all":p.param["zero_terms_query"]="all"
		}
	}
}

//default 0
func (m MatchPhrasePrefix) WithSlop(slop int) func (*MatchPhrasePrefixParam){
	return func(p *MatchPhrasePrefixParam) {
		if slop > 0{
			p.param["slop"]=slop
		}
	}
}

func (m MatchPhrasePrefix)WithName(name string) func(*MatchPhrasePrefixParam){
	return func(p *MatchPhrasePrefixParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

type MatchPhrasePrefixParam struct {
	field string
	param map[string]interface{}
}

func (m *MatchPhrasePrefixParam) Build() map[string]interface{}{
	return map[string]interface{}{
		"match_phrase_prefix":map[string]interface{}{
			m.field:m.param,
		},
	}
}
