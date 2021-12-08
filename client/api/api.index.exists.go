package api

func newIndexExistsFunc(transport Transport) IndexExists {
	return func(o ...func(*IndexExistsRequest)) (*Response, error) {
		r := IndexExistsRequest{}
		r.method=MethodHeadFunc
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type IndexExists func(o ...func(*IndexExistsRequest)) (*Response, error)

func (f IndexExists) WithIndex(v string) func(*IndexExistsRequest) {
	return func(r *IndexExistsRequest) {
		r.Index = v
	}
}

type IndexExistsRequest struct {
	BaseRequest
}
