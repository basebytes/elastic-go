package api

import (
	"strconv"
	"strings"
)

func newSearchFunc(t Transport) Search {
	return func(o ...func(*SearchRequest)) (*Response, error) {
		var r = SearchRequest{}
		for _, f := range o {
			f(&r)
		}
		r.uris=r.getUris
		return r.Do(r.ctx, t)
	}
}

type Search func(o ...func(*SearchRequest)) (*Response, error)

func (f Search) WithIndex(v ...string) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Index = strings.Join(v, ",")
	}
}

func (f Search) WithBody(v *[]byte) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Body = v
	}
}

func (f Search) WithSource(v ...string) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.Source = v
		r.AddParam("_source", strings.Join(v, ","))
	}
}

func (f Search) WithSourceExcludes(v ...string) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.SourceIncludes = v
		r.AddParam("_source_excludes", strings.Join(v, ","))
	}
}

func (f Search) WithSourceIncludes(v ...string) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.SourceIncludes = v
		r.AddParam("_source_includes", strings.Join(v, ","))
	}
}
func (f Search) WithTrackTotalHits(v bool) func(*SearchRequest) {
	return func(r *SearchRequest) {
		r.TrackTotalHits = v
		r.AddParam("track_total_hits", strconv.FormatBool(v))
	}
}

type SearchRequest struct {
	BaseRequest
	Source         []string
	SourceExcludes []string
	SourceIncludes []string
	TrackTotalHits bool
}

func (r *SearchRequest) getUris() []string {
	var uris []string
	if len(r.Index) > 0 {
		uris = append(uris, r.Index)
	}
	uris = append(uris, "_search")
	return uris
}
