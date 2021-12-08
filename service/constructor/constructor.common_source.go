package constructor

//TODO not finished
func NewSource() Source{
	return func(o ...func(*SourceParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"_source":"true",
			}
		}
		s:=&SourceParam{includes:make([]string,0),excludes:make([]string,0)}
		for _,f:=range o{
			f(s)
		}
		return  s.Build()
	}

}

type Source func(o ...func(*SourceParam)) map[string]interface{}

func(s Source)WithIncludes(fields ...string) func(*SourceParam){
	return func(p *SourceParam) {
		p.includes=append(p.includes,fields...)
	}
}

func(s Source)WithExcludes(fields ...string) func(*SourceParam){
	return func(p *SourceParam) {
		p.excludes=append(p.excludes,fields...)
	}
}

type SourceParam struct{
	includes,excludes []string
}

func (s *SourceParam)Build() map[string]interface{}{
	fields:=make(map[string]interface{},2)
	if len(s.includes)>0{
		fields["includes"]=s.includes
	}
	if len(s.excludes)>0{
		fields["excludes"]=s.excludes
	}
	if len(fields)==0{
		return map[string]interface{}{
			"_source":true,
		}
	}
	return map[string]interface{}{
		"_source":fields,
	}
}