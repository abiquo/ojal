package core

import "net/http"

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
		http.StatusNoContent: true,
	},
	http.MethodGet: map[int]bool{
		http.StatusOK: true,
	},
}

var collections = map[string]func() Resource{}

// RegisterCollection sets the Resource factory for the media collection
func RegisterCollection(media string, factory func() Resource) {
	collections[Media(media)] = factory
}

var resources = map[string]func() Resource{}

// RegisterResource sets the Resource factory for the media collection
func RegisterResource(media string, factory func() Resource) {
	resources[Media(media)] = factory
}
