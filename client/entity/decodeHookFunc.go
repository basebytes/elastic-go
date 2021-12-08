package entity

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
)

func BucketsMapToSliceHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		k := reflect.ValueOf("key")
		if f.Kind() == reflect.Map && t.Kind() == reflect.Slice && t == reflect.TypeOf([]*BucketItem{}) {
			dataVal := reflect.ValueOf(data)
			newData := make([]interface{}, dataVal.Len())
			for i, key := range dataVal.MapKeys() {
				value := dataVal.MapIndex(key)
				value.Elem().SetMapIndex(k, key)
				newData[i] = value.Interface()
			}
			return &newData, nil
		}
		return data, nil
	}
}


func AnyToStringHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		dataVal := reflect.ValueOf(data)
		switch t.Kind() {
		case reflect.String:
			switch f.Kind() {
			case reflect.Bool:
				return strconv.FormatBool(dataVal.Bool()), nil
			case reflect.Float32:
				return strconv.FormatFloat(dataVal.Float(), 'f', -1, 32), nil
			case reflect.Float64:
				return strconv.FormatFloat(dataVal.Float(), 'f', -1, 64), nil
			case reflect.Int:
				return strconv.FormatInt(dataVal.Int(), 10), nil
			case reflect.Uint:
				return strconv.FormatUint(dataVal.Uint(), 10), nil
			}
		case reflect.Struct:
			if _,e:=t.FieldByName("EsError");e{
				v:=data.(map[string]interface{})
				if err,OK:=v["error"].(map[string]interface{});OK{
					v["error"]= fmt.Sprintf(`%s:%s`,err["type"],err["reason"])
				}
			}
		}
		return data, nil
	}
}

func SettingItemHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() == reflect.Map && t == reflect.TypeOf(SettingItem{}) {
			settings := data.(map[string]interface{})
			dataVal := settings["settings"].(map[string]interface{})
			if val, OK := dataVal["index"]; OK {
				settings["settings"] = getFlatSetting("index", val)
			}
		}
		return data, nil
	}
}

func getFlatSetting(k string, v interface{}) map[string]interface{} {
	switch v.(type) {
	case string:
		return map[string]interface{}{
			k: v.(string),
		}
	case map[string]interface{}:
		result := make(map[string]interface{})
		for sk, sv := range v.(map[string]interface{}) {
			for nk, nv := range getFlatSetting(sk, sv) {
				result[k+"."+nk] = nv
			}
		}
		return result
	}
	return nil
}
