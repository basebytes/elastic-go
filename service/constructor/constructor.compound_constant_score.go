package constructor

func newConstantScore() ConstantScore{
	return func(filter map[string]interface{}, o ...func(*ConstantScoreParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"constant_score":map[string]interface{}{
					"filter":filter,
				},
			}
		}
		c:=&ConstantScoreParam{param: map[string]interface{}{
			"filter":filter,
		}}
		for _,f:=range o{
			f(c)
		}
		return c.Build()
	}
}



type ConstantScore func(filter map[string]interface{},o ...func(*ConstantScoreParam)) map[string]interface{}

//default 1.0
func (c ConstantScore) WithBoost(boost float32) func (*ConstantScoreParam){
	return func(p *ConstantScoreParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (c ConstantScore)WithName(name string) func(*ConstantScoreParam){
	return func(p *ConstantScoreParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

type ConstantScoreParam struct {
	param map[string]interface{}
}

func (c *ConstantScoreParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"constant_score":c.param,
	}
}