package entity


//mapping
type IndexMapping struct {
	EsError	`mapstructure:",squash"`
	Other map[string]*MappingItem `json:"result" mapstructure:",remain"`
}

type MappingItem struct {
	Mappings *struct {
		Dynamic    bool `json:"dynamic"`
		Properties map[string]*Property	`json:"properties,omitempty"`
	} `json:"mappings"`
}

type Property struct {
	Type       string	`json:"type,omitempty" `
	DocValues  bool		`json:"doc_values,omitempty" mapstructure:"doc_values"`
	Properties map[string]*Property `json:"properties,omitempty"`
}

func (mappings *IndexMapping) GetFields() map[string][]string {
	indexFields := map[string][]string{}
	for k, v := range mappings.Other {
		var fields []string
		for r, f := range v.Mappings.Properties {
			fields = append(fields, extractFieldName(r, f)...)
		}
		indexFields[k] = fields
	}
	return indexFields
}

func extractFieldName(cur string, property *Property) []string {
	if len(property.Properties) == 0 {
		return []string{cur}
	} else {
		var fields []string
		for k, v := range property.Properties {
			for _, f := range extractFieldName(k, v) {
				fields = append(fields, cur+"."+f)
			}
		}
		return fields
	}
}
