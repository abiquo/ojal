package core

import (
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

// Oauth contains the information for an Abiquo API oauth client
type Oauth struct {
	Token       string
	TokenSecret string
	APIKey      string
	APISecret   string
}

type oauth struct {
	client *http.Client
}

func (o Oauth) newClient() requester {
	config := oauth1.NewConfig(o.APIKey, url.QueryEscape(o.APISecret))
	token := oauth1.NewToken(o.Token, url.QueryEscape(o.TokenSecret))
	return &oauth{config.Client(oauth1.NoContext, token)}
}

func (o *oauth) do(request *http.Request) (*http.Response, error) {
	return o.client.Do(request)
}
