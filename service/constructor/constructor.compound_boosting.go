package constructor

func newBoosting()Boosting{
	return func(positive, negative map[string]interface{}, negativeBoost float32,o ...func(*BoostingParam)) map[string]interface{} {
		param:=map[string]interface{}{
			"positive":positive,
			"negative":negative,
			"negative_boost":negativeBoost,
		}
		if len(o)==0{
			return map[string]interface{}{
				"boosting":param,
			}
		}
		b:=&BoostingParam{param: param}
		for _,f:=range o{
			f(b)
		}
		return b.Build()
	}
}


type Boosting func(positive,negative map[string]interface{},negativeBoost float32,o ...func(*BoostingParam)) map[string]interface{}

//default 1.0
func (b Boosting) WithBoost(boost float32) func (*BoostingParam){
	return func(p *BoostingParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (b Boosting)WithName(name string) func(*BoostingParam){
	return func(p *BoostingParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}
type BoostingParam struct {
	param map[string]interface{}
}

func (b *BoostingParam)Build()map[string]interface{}{
	return map[string]interface{}{
		"boosting":b.param,
	}
}
