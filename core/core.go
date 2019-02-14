package core

import (
	"fmt"
	"net/http"
)

var validCodes = map[string]map[int]bool{
	http.MethodDelete: map[int]bool{
		http.StatusAccepted:  true,
		http.StatusNoContent: true,
	},
	http.MethodPost: map[int]bool{
		http.StatusCreated:   true,
		http.StatusAccepted:  true,
		http.StatusNoContent: true,
	},
	http.MethodPut: map[int]bool{
		http.StatusOK:        true,
		http.StatusCreated:   true,
		http.StatusAccepted:  true,
		http.StatusNoContent: true,
	},
	http.MethodGet: map[int]bool{
		http.StatusOK: true,
	},
}

var (
	collections = map[string]string{}
	resources   = map[string]func() Resource{}
)

// RegisterMedia sets the Resource factory for the media collection
func RegisterMedia(media, collection string, factory func() Resource) {
	if len(collection) > 0 {
		collections[getMedia(collection)] = getMedia(media)
	}
	resources[getMedia(media)] = factory
}

// Factory returns a resource of the specified media type
func Factory(media string) Resource {
	if factory := resources[getMedia(media)]; factory != nil {
		return factory()
	}
	panic(fmt.Errorf("unregistered media %q", media))
}
