package query

func NewRange() Range {
	return func(field string, o ...func(*RangeParam)) map[string]interface{} {
		if len(o) == 0 {
			return map[string]interface{}{
				"range": map[string]interface{}{
					field: struct{}{},
				},
			}
		}
		r := &RangeParam{field: field, param: map[string]interface{}{}}
		for _, f := range o {
			f(r)
		}
		return r.Build()
	}
}

type Range func(field string, o ...func(*RangeParam)) map[string]interface{}

//default Intersects
func (r Range) WithRelation(relation RelationType) func(*RangeParam) {
	return func(p *RangeParam) {
		switch relation {
		case Contains, Within:
			p.param["relation"] = relation
		}
	}
}

func (r Range) WithTimeZone(timeZone string) func(*RangeParam) {
	return func(p *RangeParam) {
		if timeZone != "" {
			p.param["time_zone"] = timeZone
		}
	}
}

func (r Range) WithDateFormat(dateFormat string) func(*RangeParam) {
	return func(p *RangeParam) {
		if dateFormat != "" {
			p.param["format"] = dateFormat
		}
	}
}

func (r Range) WithCompareOperate(operator CompareOperator, value interface{}) func(*RangeParam) {
	return func(p *RangeParam) {
		switch operator {
		case CompareOperatorGT:
			delete(p.param, "gte")
			p.param["gt"] = value
		case CompareOperatorGTE:
			delete(p.param, "gt")
			p.param["gte"] = value
		case CompareOperatorLT:
			delete(p.param, "lte")
			p.param["lt"] = value
		case CompareOperatorLTE:
			delete(p.param, "lt")
			p.param["lte"] = value
		}
	}
}

//default 1.0
func (r Range) WithBoost(boost float32) func(*RangeParam) {
	return func(p *RangeParam) {
		if boost != 1 && boost > 0 {
			p.param["boost"] = boost
		}
	}
}

func (r Range) WithName(name string) func(*RangeParam) {
	return func(p *RangeParam) {
		if name != "" {
			p.param["_name"] = name
		}
	}
}

type RangeParam struct {
	field string
	param map[string]interface{}
}

func (r *RangeParam) Build() map[string]interface{} {
	return map[string]interface{}{
		"range": map[string]interface{}{
			r.field: r.param,
		},
	}
}

type RelationType string

const (
	//Matches documents with a range field value that intersects the query’s range.
	Intersects RelationType = "INTERSECTS"
	//Matches documents with a range field value that entirely contains the query’s range.
	Contains RelationType = "CONTAINS"
	//Matches documents with a range field value entirely within the query’s range.
	Within RelationType = "WITHIN"
)

type CompareOperator string

const (
	CompareOperatorGT  CompareOperator = "gt"
	CompareOperatorGTE CompareOperator = "gte"
	CompareOperatorLT  CompareOperator = "lt"
	CompareOperatorLTE CompareOperator = "lte"
)
