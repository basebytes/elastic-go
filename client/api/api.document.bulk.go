package api

import (
	"bytes"
	"fmt"
	"github.com/basebytes/tools"
)

func newDocumentBulkFunc(transport Transport) DocumentBulk {
	return func(docs []Doc, o ...func(*DocumentBulkRequest)) (*Response, error) {
		r := DocumentBulkRequest{BaseRequest{
			method: MethodPostFunc,
		}}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		var buffer bytes.Buffer
		for _,doc:=range docs{
			if doc.BulkAction().Delete!=nil{
				buffer.WriteString(fmt.Sprintf("%s\n",tools.EncodeBytes(doc.BulkAction())))
			}else{
				buffer.WriteString(fmt.Sprintf("%s\n%s\n",tools.EncodeBytes(doc.BulkAction()),doc.Content()))
			}
		}
		body:=buffer.Bytes()
		r.Body=&body
		return r.Do(r.ctx, transport)
	}
}

type DocumentBulk func(docs []Doc, o ...func(*DocumentBulkRequest)) (*Response, error)

func (f DocumentBulk) WithIndex(v string) func(*DocumentBulkRequest) {
	return func(r *DocumentBulkRequest) {
		if v!=""{
			r.Index = v
		}
	}
}

type DocumentBulkRequest struct {
	BaseRequest
}

func (f *DocumentBulkRequest) getUris() []string {
	var uris []string
	if f.Index != "" {
		uris = append(uris, f.Index)
	}
	uris = append(uris, "_bulk")
	return uris
}
