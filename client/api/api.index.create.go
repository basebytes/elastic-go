package api

func newIndexCreteFunc(transport Transport) IndexCreate {
	return func(index string, o ...func(request *IndexCreateRequest)) (*Response, error) {
		r := IndexCreateRequest{BaseRequest: BaseRequest{
			Index: index,
			method: MethodPutFunc,
			},
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type IndexCreate func(index string, o ...func(request *IndexCreateRequest)) (*Response, error)

func (f IndexCreate) WithBody(v *[]byte) func(*IndexCreateRequest) {
	return func(r *IndexCreateRequest) {
		if v!=nil&&len(*v)>0{
			r.Body = v
		}
	}
}

type IndexCreateRequest struct {
	BaseRequest
}