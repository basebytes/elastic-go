package constructor

//TODO not finished
func NewSort() Sort{
	return func(o ...func(*SortParam)) map[string]interface{} {
		if len(o)==0{
			return map[string]interface{}{
				"sort":map[string]interface{}{
					"_score":"desc",
				},
			}
		}

		s:=&SortParam{orders:make([]map[string]interface{},0)}
		for _,f:=range o{
			f(s)
		}
		return  s.Build()
	}

}

type Sort func(o ...func(*SortParam)) map[string]interface{}

func(s Sort)WithOrder(field string,order interface{}) func(*SortParam){
	return func(p *SortParam) {
		p.orders=append(p.orders,map[string]interface{}{field:order})
	}
}

type SortParam struct{
	orders []map[string]interface{}
}

func (s *SortParam)Build() map[string]interface{}{
	return map[string]interface{}{
		"sort":s.orders,
	}
}