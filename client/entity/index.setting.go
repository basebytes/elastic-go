package entity

import (
	"regexp"
	"strings"
)

//setting
type IndexSetting struct{
	EsError	`mapstructure:",squash"`
	Other map[string]*SettingItem `json:"result" mapstructure:",remain"`
}

type SettingItem struct {
	Settings map[string]interface{} `json:"settings"`
}

func (s *IndexSetting) GetSettings(setting string) map[string]map[string]string{
	if len(s.Other) == 0 {
		return nil
	}
	settings := make(map[string]map[string]string)
	containStar:=strings.Index(setting,"*")!=-1
	if !containStar{
		setting+="*"
	}
	reg:=regexp.MustCompile(setting)
	for index, value := range s.Other {
		sets := make(map[string]string)
		for k, v := range value.Settings {
			if reg.MatchString(k) {
				sets[k] = v.(string)
			}
		}
		settings[index] = sets
	}
	return settings
}