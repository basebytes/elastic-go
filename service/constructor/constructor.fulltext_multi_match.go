package constructor

import "strings"

func newMultiMatch() MultiMatch{
	return func(text string, o ...func(*MultiMatchParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"multi_match":map[string]interface{}{
					"query":text,
				},
			}
		}
		m:=&MultiMatchParam{param: map[string]interface{}{
			"query":text,
			"type":defaultMultiMatchType,
		}}
		for _,f:=range o{
			f(m)
		}
		return m.Build()
	}
}



type MultiMatch func(text string,o ... func(*MultiMatchParam))map[string]interface{}


//default index.query.default_field or *.*
func (m MultiMatch)WithFields(fields ...string) func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if len(fields)>0{
			p.param["fields"]=fields
		}
	}
}

//default best_fields
func (m MultiMatch)WithType(matchType MultiMatchType) func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		switch matchType {
		case MultiMatchTypeMostFields,MultiMatchTypeCrossFields,
			MultiMatchTypePhrase,MultiMatchTypePhrasePrefix,MultiMatchTypeBoolPrefix:
				p.param["type"]=matchType
		}
	}
}

//default 0
func (m MultiMatch)WithTieBreaker(tieBreaker float32) func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if tieBreaker>0{
			if tieBreaker>1{
				tieBreaker=1
			}
			p.param["tie_breaker"]=tieBreaker
		}
	}
}

//default  analysis.analyzer.default or standard analyzer
func (m MultiMatch) WithAnalyzer(analyzer string)func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if analyzer!=""{
			p.param["analyzer"]=analyzer
		}
	}
}
//default 1.0
func (m MultiMatch) WithBoost(boost float32) func (*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

//default OR
func (m MultiMatch) WithOperator(operator string)func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		switch strings.ToUpper(operator) {
		case "AND":
			switch p.param["type"] {
			case MultiMatchTypePhrase,MultiMatchTypePhrasePrefix:
			default:
				p.param["operator"]="AND"
			}
		}
	}
}

func (m MultiMatch) WithMinShouldMatch(minimumShouldMatch string)func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if minimumShouldMatch!=""{
			switch p.param["type"] {
			case MultiMatchTypePhrase,MultiMatchTypePhrasePrefix:
			default:
				p.param["minimum_should_match"]=minimumShouldMatch
			}
		}
	}
}

//default none
func (m MultiMatch) WithZeroTermsQuery(zeroTermsQuery string)func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		switch strings.ToLower(zeroTermsQuery) {
		case "all":p.param["zero_terms_query"]="all"
		}
	}
}

func (m MultiMatch)WithName(name string) func(*MultiMatchParam){
	return func(p *MultiMatchParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

//best_fields
//  fuzziness, lenient, prefix_length, max_expansions, fuzzy_rewrite,  cutoff_frequency, auto_generate_synonyms_phrase_query, fuzzy_transpositions
//most_fields
//  fuzziness, lenient, prefix_length, max_expansions, fuzzy_rewrite,  cutoff_frequency
//cross_fields
//  lenient,  cutoff_frequency
//phrase
// lenient, slop
//phrase_prefix
// lenient, slop,max_expansions
//bool_prefix
//  lenient, auto_generate_synonyms_phrase_query,fuzziness, prefix_length, max_expansions, fuzzy_rewrite,fuzzy_transpositions

type MultiMatchParam struct {
	param map[string]interface{}
}

//{
//    "multi_match" : {
//      "query": <text>,
//      "fields":<fields>
//		...
//    }
//}
func (m *MultiMatchParam) Build()map[string]interface{}{
	return map[string]interface{}{
		"multi_match":m.param,
	}
}





type MultiMatchType string

const (
	defaultMultiMatchType = MultiMatchTypeBestFields
	//Finds documents which match any field, but uses the _score from the best field.
	MultiMatchTypeBestFields MultiMatchType="best_fields"
	//Finds documents which match any field and combines the _score from each field
	MultiMatchTypeMostFields MultiMatchType="most_fields"
	//Treats fields with the same analyzer as though they were one big field. Looks for each word in any field
	MultiMatchTypeCrossFields MultiMatchType="cross_fields"
	//Runs a match_phrase query on each field and uses the _score from the best field.
	MultiMatchTypePhrase MultiMatchType="phrase"
	//Runs a match_phrase_prefix query on each field and uses the _score from the best field
	MultiMatchTypePhrasePrefix MultiMatchType="phrase_prefix"
	//Creates a match_bool_prefix query on each field and combines the _score from each field.
	MultiMatchTypeBoolPrefix MultiMatchType="bool_prefix"
)