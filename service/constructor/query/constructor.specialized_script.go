package query

func NewScript() Script {
	return func(script func(*ScriptParam), o ...func(*ScriptParam)) map[string]interface{} {
		p := &ScriptParam{param: make(map[string]interface{})}
		o = append(o, script)
		for _, f := range o {
			f(p)
		}
		return p.Build()
	}
}

type Script func(script func(*ScriptParam), o ...func(*ScriptParam)) map[string]interface{}

func (s Script) WithLang(lang string) func(*ScriptParam) {
	return func(p *ScriptParam) {
		if lang != "" {
			p.param["lang"] = lang
		}
	}
}

func (s Script) WithScriptSource(source string, params map[string]interface{}) func(*ScriptParam) {
	return func(p *ScriptParam) {
		if source != "" {
			p.param["source"] = source
			delete(p.param, "id")
			if params != nil {
				p.param["params"] = params
			}
		}
	}
}

func (s Script) WithScriptId(id string, params map[string]interface{}) func(*ScriptParam) {
	return func(p *ScriptParam) {
		if id != "" {
			p.param["id"] = id
			delete(p.param, "source")
			if params != nil {
				p.param["params"] = params
			}
		}
	}
}

type ScriptParam struct {
	param map[string]interface{}
}

func (s *ScriptParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"script": map[string]interface{}{
			"script": s.param,
		},
	}
}
