package core

import (
	"encoding/json"
	"fmt"
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
	if res != nil {
		r = &Reply{Response: res, payload: payload}
		r.result, err = ioutil.ReadAll(r.Body)
		res.Body.Close()
		if err == nil {
			err = r.error()
		}
	}
	debug(r, err)
	return
}

// Ok returns if the reply was successful or not
func (r *Reply) Ok() (ok bool) {
	return r != nil && r.Response != nil && validCodes[r.Request.Method][r.StatusCode]
}

// Location returns an http.Response location values
func (r *Reply) Location() (location string) {
	if r != nil {
		if href, err := r.Response.Location(); err == nil {
			location = href.String()
		}
	}
	return
}

// Status returns the reply http reponse status string
func (r *Reply) Status() (status int) {
	if r != nil && r.Response != nil {
		status = r.StatusCode
	}
	return
}

func (r *Reply) error() (err error) {
	if !r.Ok() {
		msg := fmt.Sprint(r.Status(), " Unexpected status code")
		err = &Errors{status: r.StatusCode, msg: msg}
		json.Unmarshal(r.result, err)
	}
	return
}
