package api

func newIndexGetMappingFunc(transport Transport) IndexGetMapping {
	return func(o ...func(request *IndexGetMappingRequest)) (*Response, error) {
		r := IndexGetMappingRequest{}
		r.uris=r.getUris
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, transport)
	}
}

type IndexGetMapping func(o ...func(request *IndexGetMappingRequest)) (*Response, error)

func (f IndexGetMapping) WithIndex(v string) func(*IndexGetMappingRequest) {
	return func(r *IndexGetMappingRequest) {
		r.Index = v
	}
}

type IndexGetMappingRequest struct {
	BaseRequest
}

func (r *IndexGetMappingRequest) getUris() []string {
	var uris []string
	if r.Index != "" {
		uris = append(uris, r.Index)
	}
	uris = append(uris, "_mapping")
	return uris
}
