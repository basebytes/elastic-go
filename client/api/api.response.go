package api

import (
	"encoding/json"
	"fmt"
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type Response struct {
	*http.Response
	rspData []byte
	req *http.Request
}

func (r *Response) IsError() bool {
	return r.StatusCode > 299
}

func (r *Response) ResultBytes() []byte {
	return r.rspData
}

func (r *Response) ResultJson(v interface{}) error{
	if data:=r.ResultBytes();data!=nil{
		return json.Unmarshal(data, &v)
	}
	return nil
}

func (r *Response) ResultString() string {
	if data:=r.ResultBytes();data!=nil{
		return string(data)
	}
	return ""
}

func (r *Response) Result(result interface{}) (err error){
	//var (
	//	esError *entity.EsError
	//	err error
	//)
	//if r.IsError() {
	//	//fmt.Println(r.ResultString())
	//	decoderConfig.Result = result
	//	decoder, _ := mapstructure.NewDecoder(decoderConfig)
	//	decoder.Decode(res)
	//	//attempt parse to esError
	//	if err = r.ResultJson(&esError); err == nil {
	//		//err = errors.New(esError.Error.ErrType + "->" + esError.Error.Reason)
	//		//log.Printf( "[%s] %s: %s", r.Status, esError.Error.ErrType, esError.Error.Reason)
	//		return err
	//	}
	//}
	var res map[string]interface{}
	if err=r.ResultJson(&res); err == nil {
		decoderConfig.Result = result
		decoder, _ := mapstructure.NewDecoder(decoderConfig)
		err=decoder.Decode(res)
	}else{
		err=fmt.Errorf("Error parsing the response body: %s ", err)
	}
	return err
}


var decoderConfig = &mapstructure.DecoderConfig{
	WeaklyTypedInput:true,
	DecodeHook: mapstructure.ComposeDecodeHookFunc(
		entity.BucketsMapToSliceHookFunc(),
		entity.AnyToStringHookFunc(),
		entity.SettingItemHookFunc()),
}
