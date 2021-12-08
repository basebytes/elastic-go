package constructor

import "strings"

func newCombinedFields() CombinedFields{
	return func(fields []string, text string, o ...func(*CombinedFieldsParam)) map[string]interface{} {
		c:=&CombinedFieldsParam{
			param: map[string]interface{}{
				"query":text,
				"fields":fields,
			},
		}
		for _,f:=range o{
			f(c)
		}
		return c.Build()
	}
}

type CombinedFields func(fields []string,text string,o ...func(*CombinedFieldsParam)) map[string]interface{}


//default OR
func (c CombinedFields) WithOperator(operator string)func(*CombinedFieldsParam){
	return func(p *CombinedFieldsParam) {
		switch strings.ToUpper(operator) {
		case "AND":p.param["operator"]="AND"
		}
	}
}

func (c CombinedFields) WithMinShouldMatch(minimumShouldMatch string)func(*CombinedFieldsParam){
	return func(p *CombinedFieldsParam) {
		if minimumShouldMatch!=""{
			p.param["minimum_should_match"]=minimumShouldMatch
		}
	}
}

//default none
func (c CombinedFields) WithZeroTermsQuery(zeroTermsQuery string)func(*CombinedFieldsParam){
	return func(p *CombinedFieldsParam) {
		switch strings.ToLower(zeroTermsQuery) {
		case "all":p.param["zero_terms_query"]="all"
		}
	}
}
//default 1.0
func (c CombinedFields) WithBoost(boost float32) func (*CombinedFieldsParam){
	return func(p *CombinedFieldsParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (c CombinedFields)WithName(name string) func(*CombinedFieldsParam){
	return func(p *CombinedFieldsParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}
//auto_generate_synonyms_phrase_query

type CombinedFieldsParam struct{
	param map[string]interface{}
}

func (c *CombinedFieldsParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"combined_fields":c.param,
	}
}