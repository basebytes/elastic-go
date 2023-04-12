package query

func NewTerms() Terms {
	return func(field string, o ...func(*TermsParam)) map[string]interface{} {
		p := &TermsParam{field: field, param: map[string]interface{}{}}
		for _, f := range o {
			f(p)
		}
		return p.Build()
	}
}

//{
//	"terms":{
//		<field>:[<value>...],
//      "boost":<boost>,
//	}
//}
//
//or
//
//{
//	"terms":{
//		<field>:{
//			"index":<index>,
//			"id":<id>,
//			"path":<path>,
//			"routing":<routing>
//		},
//		"boost":<boost>
//	}
//}
type Terms func(field string, o ...func(*TermsParam)) map[string]interface{}

func (t Terms) WithValue(value ...interface{}) func(*TermsParam) {
	return func(p *TermsParam) {
		if len(value) > 0 {
			p.param[p.field] = value
		}
	}
}

//default 1.0
func (t Terms) WithBoost(boost float32) func(*TermsParam) {
	return func(p *TermsParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

func (t Terms) WithLookUp(index, id, path, routing string) func(*TermsParam) {
	return func(p *TermsParam) {
		lookUp := map[string]string{
			"index": index,
			"id":    id,
			"path":  path,
		}
		if routing != "" {
			lookUp["routing"] = routing
		}
		p.param[p.field] = lookUp
	}
}

func (t Terms) WithName(name string) func(*TermsParam) {
	return func(p *TermsParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

type TermsParam struct {
	field string
	param map[string]interface{}
}

func (t *TermsParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"terms": t.param,
	}
}
