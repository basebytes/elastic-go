package query

import "strings"

func NewSimpleQueryString() SimpleQueryString {
	return func(text string, o ...func(*SimpleQueryStringParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"simple_query_string": map[string]interface{}{
					"query": text,
				},
			}
		}
		s := &SimpleQueryStringParam{param: map[string]interface{}{
			"query": text,
		}}
		for _, f := range o {
			f(s)
		}
		return s.Build()
	}
}

//default  index.query.default_field  or *
func (s SimpleQueryString) WithFields(fields ...string) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if len(fields) > 0 {
			p.param["fields"] = fields
		}
	}
}

//default  analysis.analyzer.default or standard analyzer
func (s SimpleQueryString) WithAnalyzer(analyzer string) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if analyzer != "" {
			p.param["analyzer"] = analyzer
		}
	}
}

//default false
func (s SimpleQueryString) WithAnalyzerWildcard(analyzerWildcard bool) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if analyzerWildcard {
			p.param["analyze_wildcard"] = analyzerWildcard
		}
	}
}

//default OR
func (s SimpleQueryString) WithDefaultOperator(defaultOperator string) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		switch strings.ToUpper(defaultOperator) {
		case "AND":
			p.param["default_operator"] = "AND"
		}
	}
}

func (s SimpleQueryString) WithMinShouldMatch(minimumShouldMatch string) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if minimumShouldMatch != "" {
			p.param["minimum_should_match"] = minimumShouldMatch
		}
	}
}

//default All
func (s SimpleQueryString) WithFlags(flags ...Flag) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		var fs []string
		for _, f := range flags {
			switch f {
			case And, Escape, Fuzzy, Near, None, Not, Or, Phrase, Precedence, Prefix, Slop, Whitespace:
				fs = append(fs, string(f))
			}
		}
		if len(fs) > 0 {
			p.param["flags"] = strings.Join(fs, "|")
		}
	}
}

//default 1.0
func (s SimpleQueryString) WithBoost(boost float32) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

func (s SimpleQueryString) WithName(name string) func(*SimpleQueryStringParam) {
	return func(p *SimpleQueryStringParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

//auto_generate_synonyms_phrase_query
//fuzzy_max_expansions
//fuzzy_prefix_length
//fuzzy_transpositions
//lenient
//quote_field_suffix

type SimpleQueryString func(text string, o ...func(*SimpleQueryStringParam)) map[string]interface{}

type SimpleQueryStringParam struct {
	param map[string]interface{}
}

func (s *SimpleQueryStringParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"simple_query_string": s.param,
	}
}

type Flag string

const (
	//Enables the + AND operator.
	And Flag = "AND"
	//Enables \ as an escape character.
	Escape Flag = "ESCAPE"
	//Enables the ~N operator after a word, where N is an integer denoting the allowed edit distance for matching. See Fuzziness.
	Fuzzy Flag = "FUZZY"
	//Enables the ~N operator, after a phrase where N is the maximum number of positions allowed between matching tokens. Synonymous to SLOP.
	Near Flag = "NEAR"
	//Disables all operators.
	None Flag = "NONE"
	//Enables the - NOT operator.
	Not Flag = "NOT"
	//Enables the \| OR operator.
	Or Flag = "OR"
	//Enables the " quotes operator used to search for phrases.
	Phrase Flag = "PHRASE"
	//Enables the ( and ) operators to control operator precedence.
	Precedence Flag = "PRECEDENCE"
	//Enables the * prefix operator.
	Prefix Flag = "PREFIX"
	//Enables the ~N operator, after a phrase where N is maximum number of positions allowed between matching tokens. Synonymous to NEAR.
	Slop Flag = "SLOP"
	//Enables whitespace as split characters.
	Whitespace Flag = "WHITESPACE"
)
