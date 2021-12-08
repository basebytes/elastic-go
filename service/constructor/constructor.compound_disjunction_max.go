package constructor

func newDisjunctionMax() DisjunctionMax{
	return func(queries []map[string]interface{}, o ...func(*DisjunctionMaxParam)) map[string]interface{} {
		param:=map[string]interface{}{
			"queries":queries,
		}
		if len(o)==0{
			return map[string]interface{}{
				"dis_max":param,
			}
		}
		d:=&DisjunctionMaxParam{param: param}
		for _,f:=range o{
			f(d)
		}
		return d.Build()
	}
}


type DisjunctionMax func(queries []map[string]interface{},o ...func(*DisjunctionMaxParam) )map[string]interface{}

//default 0
func (d DisjunctionMax)WithTieBreaker(tieBreaker float32) func(*DisjunctionMaxParam){
	return func(p *DisjunctionMaxParam) {
		if tieBreaker>0{
			if tieBreaker>1{
				tieBreaker=1
			}
			p.param["tie_breaker"]=tieBreaker
		}
	}
}

//default 1.0
func (d DisjunctionMax) WithBoost(boost float32) func (*DisjunctionMaxParam){
	return func(p *DisjunctionMaxParam) {
		if boost != 1 && boost > 0{
			p.param["boost"]=boost
		}
	}
}

func (d DisjunctionMax)WithName(name string) func(*DisjunctionMaxParam){
	return func(p *DisjunctionMaxParam) {
		if name!=""{
			p.param["_name"]=name
		}
	}
}

type DisjunctionMaxParam struct {
	param map[string]interface{}
}

func (d *DisjunctionMaxParam)Build()map[string]interface{}{
	return map[string]interface{}{
		"dis_max":d.param,
	}
}