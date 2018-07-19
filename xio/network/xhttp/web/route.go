package web

import (
	"net/http"
	"strings"
)

type routeCtxKey int

var routeKey routeCtxKey = 1

type route struct {
	path string
	last string
}

func (r *route) shift() string {
	i := strings.Index(r.path[1:], "/") + 1
	if i <= 0 {
		r.last = r.path[1:]
		r.path = "/"
	} else {
		r.last = r.path[1:i]
		r.path = r.path[i:]
	}
	return r.last
}

func (r *route) remaining() string {
	return r.path[1:]
}

// PathHeadThenShift returns the head segment of the request's adjusted path,
// then shifts it left by one segment. This does not adjust the path stored in
// req.URL.Path.
func PathHeadThenShift(req *http.Request) string {
	r, ok := req.Context().Value(routeKey).(*route)
	if !ok {
		return req.URL.Path
	}
	return r.shift()
}

// LastPathHead returns the last result obtained from a call to
// PathHeadThenShift() for the request.
func LastPathHead(req *http.Request) string {
	r, ok := req.Context().Value(routeKey).(*route)
	if !ok {
		return req.URL.Path
	}
	return r.last
}

// HasMorePathSegments returns true if more path segments will be returned
// from future calls to PathHeadThenShift() for the request.
func HasMorePathSegments(req *http.Request) bool {
	r, ok := req.Context().Value(routeKey).(*route)
	if !ok {
		return false
	}
	return r.path != "/"
}

// RemainingPath returns the remaining path for a request.
func RemainingPath(req *http.Request) string {
	r, ok := req.Context().Value(routeKey).(*route)
	if !ok {
		return req.URL.Path
	}
	return r.remaining()
}
