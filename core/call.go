package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Call represents a generic API call
type Call struct {
	method  string
	href    Href
	query   url.Values
	accept  Type
	content Type
	payload []byte
}

func newCall(method string, href Href, accept, content Type, payload []byte) *Call {
	return &Call{
		method:  method,
		href:    href,
		accept:  accept,
		content: content,
		payload: payload,
	}
}

// Delete returns a new DELETE call
func Delete(href Href) *Call {
	return newCall("DELETE", href, nil, nil, nil)
}

// Get returns a new GET call
func Get(href Href, media Type) *Call {
	return newCall("GET", href, media, nil, nil)
}

// Put returns a new PUT call
func Put(href Href, accept, content Type, payload interface{}) *Call {
	data, _ := json.Marshal(payload)
	return newCall("PUT", href, accept, content, data)
}

// Post returns a new POST call
func Post(href Href, accept, content Type, payload interface{}) *Call {
	data, _ := json.Marshal(payload)
	return newCall("POST", href, accept, content, data)
}

// Query sets an Abiquo API call query values
func (c *Call) Query(query url.Values) *Call {
	c.query = query
	return c
}

func (c *Call) request() (req *http.Request, err error) {
	reader := bytes.NewReader(c.payload)
	location := Resolve(c.href.URL(), c.query)
	req, err = http.NewRequest(c.method, location, reader)
	if err != nil {
		return
	}

	if c.accept != nil {
		req.Header.Set("Accept", c.accept.Media())
	}
	if c.content != nil {
		req.Header.Set("Content-Type", c.content.Media())
	}

	return
}
