package query

func NewIDS() IDS {
	return func(ids []string, o ...func(*IDSParam)) map[string]interface{} {
		param := map[string]interface{}{
			"values": ids,
		}
		if len(o) == 0 {
			return map[string]interface{}{
				"ids": param,
			}
		}
		i := &IDSParam{param: param}
		for _, f := range o {
			f(i)
		}
		return i.Build()
	}
}

type IDS func(ids []string, o ...func(*IDSParam)) map[string]interface{}

func (i IDS) WithName(name string) func(*IDSParam) {
	return func(p *IDSParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

//default 1.0
func (i IDS) WithBoost(boost float32) func(*IDSParam) {
	return func(p *IDSParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

type IDSParam struct {
	param map[string]interface{}
}

func (i IDSParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"ids": i.param,
	}
}
