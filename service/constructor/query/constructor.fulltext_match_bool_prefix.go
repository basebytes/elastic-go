package query

import "strings"

func NewMatchBoolPrefix() MatchBoolPrefix {
	return func(field, text string, o ...func(*MatchBoolPrefixParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"match_bool_prefix": map[string]interface{}{
					field: text,
				},
			}
		}
		m := &MatchBoolPrefixParam{field: field, param: map[string]interface{}{
			"query": text,
		}}
		for _, f := range o {
			f(m)
		}
		return m.Build()
	}
}

type MatchBoolPrefix func(field, text string, o ...func(*MatchBoolPrefixParam)) map[string]interface{}

//default  analysis.analyzer.default or standard analyzer
func (m MatchBoolPrefix) WithAnalyzer(analyzer string) func(*MatchBoolPrefixParam) {
	return func(p *MatchBoolPrefixParam) {
		if analyzer != "" {
			p.param["analyzer"] = analyzer
		}
	}
}

//default OR
func (m MatchBoolPrefix) WithOperator(operator string) func(*MatchBoolPrefixParam) {
	return func(p *MatchBoolPrefixParam) {
		switch strings.ToUpper(operator) {
		case "AND":
			p.param["operator"] = "AND"
		}
	}
}

func (m MatchBoolPrefix) WithMinShouldMatch(minimumShouldMatch string) func(*MatchBoolPrefixParam) {
	return func(p *MatchBoolPrefixParam) {
		if minimumShouldMatch != "" {
			p.param["minimum_should_match"] = minimumShouldMatch
		}
	}
}

//default 1.0
func (m MatchBoolPrefix) WithBoost(boost float32) func(*MatchBoolPrefixParam) {
	return func(p *MatchBoolPrefixParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

func (m MatchBoolPrefix) WithName(name string) func(*MatchBoolPrefixParam) {
	return func(p *MatchBoolPrefixParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

//fuzziness, prefix_length, max_expansions, fuzzy_transpositions,fuzzy_rewrite

type MatchBoolPrefixParam struct {
	field string
	param map[string]interface{}
}

func (m *MatchBoolPrefixParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"match_bool_prefix": map[string]interface{}{
			m.field: m.param,
		},
	}
}
