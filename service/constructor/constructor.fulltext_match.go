package constructor

import "strings"

func newMatch() Match{
	return func(field,text string,o ...func (*MatchParam)) map[string]interface{}{
		if len(o)==0{
			return map[string]interface{}{
				"match":map[string]interface{}{
					field:text,
				},
			}
		}

		m:=&MatchParam{field:field,param: map[string]interface{}{
			"query":text,
		}}
		for _,f:=range o{
			f(m)
		}
		return m.Build()
	}
}

type Match func(field,text string,o ...func (*MatchParam)) map[string]interface{}

//https://www.elastic.co/guide/en/elasticsearch/reference/7.13/specify-analyzer.html#specify-index-time-default-analyzer

//default  analysis.analyzer.default or standard analyzer
func (m Match) WithAnalyzer(analyzer string)func(*MatchParam){
	return func(p *MatchParam) {
		if analyzer!=""{
			p.param["analyzer"]=analyzer
		}
	}
}

//default OR
func (m Match) WithOperator(operator string)func(*MatchParam){
	return func(p *MatchParam) {
		switch strings.ToUpper(operator) {
		case "AND":p.param["operator"]="AND"
		}
	}
}

func (m Match) WithMinShouldMatch(minimumShouldMatch string)func(*MatchParam){
	return func(p *MatchParam) {
		if minimumShouldMatch!=""{
			p.param["minimum_should_match"]=minimumShouldMatch
		}
	}
}

//default none
func (m Match) WithZeroTermsQuery(zeroTermsQuery string)func(*MatchParam){
	return func(p *MatchParam) {
		switch strings.ToLower(zeroTermsQuery) {
		case "all":p.param["zero_terms_query"]="all"
		}
	}
}

//default 1.0
func (m Match) WithBoost(boost float32) func (*MatchParam){
	return func(p *MatchParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}
func (m Match)WithName(name string) func(*MatchParam){
	return func(p *MatchParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}
//func (m Match) WithOperator(operator string)func(*MatchParam){
//	return func(p *MatchParam) {
//
//	}
//}
type MatchParam struct{
	field string
	param map[string]interface{}
}

//analyzer
//operator
//minimum_should_match
//zero_terms_query
//boost
//TODO
//auto_generate_synonyms_phrase_query
//fuzziness
//max_expansions
//prefix_length
//fuzzy_transpositions
//fuzzy_rewrite
//lenient

//{
//	"match":{
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
//	"match":{
//		<field>:<text>,
//	}
//}
func (m *MatchParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"match":map[string]interface{}{
			m.field:m.param,
		},
	}
}


//func newMatch() Match{
//	return nil
//}
//
//type Match func(field string,o ...func (*MatchParam)) map[string]interface{}
//
//
//type MatchParam struct{
//	field string
//	value []string
//	boost float32
//	lookUp map[string]string
//}
//
//func (m *MatchParam)Build() map[string]interface{}{
//	return nil
//}