package core

import (
	"io/ioutil"
	"net/http"
)

// Reply represents a generic Abiquo API response
type Reply struct {
	*http.Response
	payload []byte
	result  []byte
}

func newReply(res *http.Response, payload []byte) (r *Reply, err error) {
	if res == nil {
		return
	}

	r = &Reply{Response: res, payload: payload}
	r.result, err = ioutil.ReadAll(r.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}

	if !r.Ok() {
		err = newError(res.StatusCode, r.result)
	}
	debug(r, err)
	return
}

// Ok returns if the reply was successful or not
func (r *Reply) Ok() (ok bool) {
	if r == nil {
		return
	}
	return r.Response != nil && validCodes[r.Request.Method][r.StatusCode]
}

// URL returns an http.Response location values
func (r *Reply) URL() (location string) {
	if r != nil {
		if href, err := r.Response.Location(); err == nil {
			location = href.String()
		}
	}
	return
}

// Error returns a failed reply error
func (r *Reply) Error() (e Error) {
	if !r.Ok() {
		newError(r.StatusCode, r.result)
	}
	return
}

// Media ...
func (r *Reply) Media() string {
	content, ok := r.Header["Content-Type"]
	if !ok {
		panic("no content type header")
	}

	return content[0]
}
