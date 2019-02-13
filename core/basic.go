package core

import "net/http"

// Basic contains the information for an Abiquo API basic auth client
type Basic struct {
	Username string
	Password string
}

type basic struct {
	Basic
	*http.Client
}

func (b Basic) newClient() requester {
	return &basic{b, &http.Client{}}
}

func (b *basic) do(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(b.Username, b.Password)
	return b.Client.Do(request)
}
