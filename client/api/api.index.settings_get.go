package api

import (
	"strings"
)

func newSettingsGet(transport Transport) SettingsGet {
	return func(o ...func(*SettingsGetRequest)) (*Response, error) {
		var r = SettingsGetRequest{}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type SettingsGet func(o ...func(*SettingsGetRequest)) (*Response, error)

func (f SettingsGet) WithIndex(v ...string) func(*SettingsGetRequest) {
	return func(r *SettingsGetRequest) {
		if len(v)>0{
			r.Index = strings.Join(v, ",")
		}
	}
}

func (f SettingsGet) WithSettings(v ...string) func(*SettingsGetRequest) {
	return func(r *SettingsGetRequest) {
		if len(v)>0{
			r.Settings = v
		}
	}
}

//func (f SettingsGet) WithFlatSettings(v bool) func(*SettingsGetRequest) {
//	return func(r *SettingsGetRequest) {
//		if v{
//			r.FlatSettings = v
//			r.AddParam("flat_settings", strconv.FormatBool(v))
//		}
//	}
//}

type SettingsGetRequest struct {
	BaseRequest
	Settings     []string
	FlatSettings bool
}
 func (r *SettingsGetRequest) getUris() []string {
	var uris []string
	if r.Index != "" {
		uris = append(uris, r.Index)
	}
	uris = append(uris, "_settings")
	uris = append(uris, strings.Join(r.Settings, ","))
	return uris
}
