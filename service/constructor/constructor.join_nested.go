package constructor

func newNested() Nested{
	return func(path string, query map[string]interface{}, o ...func(param *NestedParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"nested":map[string]interface{}{
					"path":path,
					"query":query,
				},
			}
		}
		n:=&NestedParam{path:path,query: query}
		for _,f:=range o{
			f(n)
		}
		return n.Build()
	}
}

type Nested func(path string,query map[string]interface{},o ...func(param *NestedParam))map[string]interface{}

func (n Nested)WithSource(source map[string]interface{}) func(*NestedParam){
	return func(p *NestedParam) {
		if len(source)>0{
			p.source=source
		}
	}
}

func (n Nested)WithReturn(whetherReturn bool) func(*NestedParam){
	return func(p *NestedParam) {
		p.whether=whetherReturn
	}
}

func (n Nested)WithName(name string) func(*NestedParam){
	return func(p *NestedParam) {
		if name!=""{
			p.name=name
		}
	}
}

type NestedParam struct{
	path,name string
	whether bool
	query,source map[string]interface{}
}

func (n *NestedParam)Build() map[string]interface{}{
	innerHits:=make(map[string]interface{})
	nested:= map[string]interface{}{
		"path":       n.path,
		"query":      n.query,
		"inner_hits": innerHits,
	}
	if n.name!=""{
		innerHits["name"]=n.name
	}

	if !n.whether{
		innerHits["_source"]=false
	}else if len(n.source)>0{
		innerHits["-_source"]=n.source["_source"]
	}

	return map[string]interface{}{
		"nested": nested,
	}
}