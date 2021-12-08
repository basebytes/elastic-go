package api


func newPingFunc(transport Transport) Ping{
	return func() (*Response, error) {
		var r = PingRequest{}
		return r.Do(r.ctx, transport)
	}
}

type Ping func() (*Response, error)


type PingRequest struct {
	BaseRequest
}

