package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dghubble/oauth1"
)

var (
	client interface {
		Do(*http.Request) (*http.Response, error)
	}
	location *url.URL
	version  string
)

// Basic contains the information for an Abiquo API basic auth client
type Basic struct {
	Username string
	Password string
}

type basic struct {
	Basic
	*http.Client
}

func (b *basic) Do(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(b.Username, b.Password)
	return b.Client.Do(request)
}

// Oauth contains the information for an Abiquo API oauth client
type Oauth struct {
	Token       string
	TokenSecret string
	APIKey      string
	APISecret   string
}

type oauth struct {
	*http.Client
}

func (o *oauth) Do(request *http.Request) (*http.Response, error) {
	return o.Client.Do(request)
}

// Init initializes Abiquo API client
func Init(api string, auth interface{}) (err error) {
	if location, err = url.Parse(api + "/"); err == nil {
		switch t := auth.(type) {
		case Basic:
			client = &basic{t, &http.Client{}}
		case Oauth:
			config := oauth1.NewConfig(t.APIKey, url.QueryEscape(t.APISecret))
			token := oauth1.NewToken(t.Token, url.QueryEscape(t.TokenSecret))
			client = &oauth{config.Client(oauth1.NoContext, token)}
		}

		var r *Reply
		if r, err = Rest(nil, Get(Resolve("version", nil), "text/plain")); r.Ok() {
			version = strings.Join(strings.Split(string(r.result), ".")[0:2], ".")
		}
	}
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

// Version returns the Abiquo API client version
func Version() string { return version }

// Do performs an Abiquo API call
func do(c *Call) (r *Reply, err error) {
	var req *http.Request
	if req, err = c.request(); err == nil {
		var res *http.Response
		if res, err = client.Do(req); err == nil {
			r, err = newReply(res, c.payload)
		}
	}
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
	} else {
		err = r.error()
	}
	return
}
