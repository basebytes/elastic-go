package constructor

func newTerm() Term{
	return func(field,value string, o ...func(*TermParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"term": map[string]interface{}{
					field: value,
				},
			}
		}
		p:=&TermParam{field:field,param: map[string]interface{}{
			"value":value,
		}}
		for _,f:=range o{
			f(p)
		}
		return p.Build()
	}
}

//{
//	"term":{
//		<field>:<value>
//	}
//}
//
//or
//
//{
//	"term":{
//		<field>:{
//			"value":<value>,
//			"boost":<boost>,
//			"case_insensitive":<caseInsensitive>
//		}
//	}
//}
type Term func(field,value string,o ...func (*TermParam)) map[string]interface{}

//default false
func (t Term)WithCaseInsensitive(caseInsensitive bool) func (*TermParam){
	return func(p *TermParam) {
		if caseInsensitive{
			p.param["case_insensitive"]=caseInsensitive
		}
	}
}

func (t Term)WithName(name string) func(*TermParam){
	return func(p *TermParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

//default 1.0
func (t Term) WithBoost(boost float32) func (*TermParam){
	return func(p *TermParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

type TermParam struct{
	field string
	param map[string]interface{}
}

func (t *TermParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"term":map[string]interface{}{
			t.field:t.param,
		},
	}
}