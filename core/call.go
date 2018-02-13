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
	href    string
	query   url.Values
	accept  string
	content string
	payload []byte
}

// Delete returns a new DELETE call
func Delete(href string) *Call {
	return &Call{method: "DELETE", href: href}
}

// Get returns a new GET call
func Get(href, media string) *Call {
	return &Call{method: "GET", href: href, accept: media}
}

// Put returns a new PUT call
func Put(href, accept, content string, payload interface{}) *Call {
	data, _ := json.Marshal(payload)
	return &Call{
		method:  "PUT",
		href:    href,
		accept:  accept,
		content: content,
		payload: data,
	}
}

// Post returns a new POST call
func Post(href, accept, content string, payload interface{}) *Call {
	data, _ := json.Marshal(payload)
	return &Call{
		method:  "POST",
		href:    href,
		accept:  accept,
		content: content,
		payload: data,
	}
}

// Query sets an Abiquo API call query values
func (c *Call) Query(query url.Values) *Call {
	c.query = query
	return c
}

func (c *Call) request() (req *http.Request, err error) {
	reader := bytes.NewReader(c.payload)
	location := Resolve(c.href, c.query)
	if req, err = http.NewRequest(c.method, location, reader); err == nil {
		if c.accept != "" {
			req.Header.Set("Accept", Media(c.accept))
		}
		if c.content != "" {
			req.Header.Set("Content-Type", Media(c.content))
		}
	}
	return
}
