package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type requester interface {
	do(*http.Request) (*http.Response, error)
}

// Authenticator ...
type Authenticator interface {
	newClient() requester
}

var (
	client   requester
	location *url.URL
)

// Init initializes Abiquo API client
func Init(api string, auth Authenticator) (err error) {
	location, err = url.Parse(api + "/")
	if err != nil {
		return
	}

	client = auth.newClient()
	return
}

const mediaFmt = "application/vnd.abiquo.%v+json"

// Media returns the equivalent media type for a media shortcut
func Media(media string) string {
	if media != "" && !strings.ContainsAny(media, "/") {
		media = fmt.Sprintf(mediaFmt, media)
	}
	return media
}

// Resolve resolves a rawurl location iwithin the api endpoint
func Resolve(href string, query url.Values) string {
	u, _ := url.Parse(href)
	q := u.Query()
	for key, values := range query {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return location.ResolveReference(u).String()
}

// Do performs an Abiquo API call
func do(c *Call) (r *Reply, err error) {
	req, err := c.request()
	if err != nil {
		return
	}

	res, err := client.do(req)
	if err != nil {
		return
	}

	r, err = newReply(res, c.payload)
	return
}

// Rest provides a facility to perform a rest Call
func Rest(result interface{}, c *Call) (r *Reply, err error) {
	r, err = do(c)
	if err != nil {
		return
	}

	if r.Ok() {
		json.Unmarshal(r.result, result)
	}
	return
}
