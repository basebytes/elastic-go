package constructor

func newMatchAll() MatchAll{
	return func(o ...func (*MatchAllParam)) map[string]interface{}{
		if len(o)==0{
			return map[string]interface{}{
				"match_all":struct{}{},
			}
		}
		m:=&MatchAllParam{param: map[string]interface{}{}}
		for _,f:=range o{
			f(m)
		}
		return m.Build()
	}
}

type MatchAll func(o ...func (*MatchAllParam)) map[string]interface{}

//default 1.0
func (m MatchAll) WithBoost(boost float32) func (*MatchAllParam){
	return func(p *MatchAllParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (m MatchAll)WithName(name string) func(*MatchAllParam){
	return func(p *MatchAllParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

type MatchAllParam struct {
	param map[string]interface{}
}

func (m *MatchAllParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"match_all": m,
	}
}