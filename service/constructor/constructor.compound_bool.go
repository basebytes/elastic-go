package constructor

func newBool() Bool{
	return func(o ...func(*BoolParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"bool": struct {}{},
			}
		}
		b:=&BoolParam{param:map[string]interface{}{}}
		for _,f:=range o{
			f(b)
		}
		return b.Build()
	}
}



type Bool func(o ...func (*BoolParam)) map[string]interface{}

func(b Bool)WithClause(clauseType ClauseType,queries ...map[string]interface{}) func(*BoolParam){
	return func(p *BoolParam) {
		if c:=len(queries);c>0{
			switch clauseType {
			case ClauseTypeShould,ClauseTypeFilter,ClauseTypeMust,ClauseTypeMustNot:
				p.param[string(clauseType)]=queries
			}
		}
	}
}

func (b Bool) WithMinShouldMatch(minimumShouldMatch string)func(*BoolParam){
	return func(p *BoolParam) {
		if minimumShouldMatch!=""{
			p.param["minimum_should_match"]=minimumShouldMatch
		}
	}
}

//default 1.0
func (b Bool) WithBoost(boost float32) func (*BoolParam){
	return func(p *BoolParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (b Bool)WithName(name string) func(*BoolParam){
	return func(p *BoolParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

type BoolParam struct {
	param map[string]interface{}
}

func (b *BoolParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"bool":b.param,
	}
}

type ClauseType string

const (
	ClauseTypeMust ClauseType="must"
	ClauseTypeMustNot ClauseType="must_not"
	ClauseTypeShould ClauseType="should"
	ClauseTypeFilter ClauseType="filter"
)