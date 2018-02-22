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

var (
	collections = map[string]func() Resource{}
	resources   = map[string]func() Resource{}
)

// RegisterMedia sets the Resource factory for the media collection
func RegisterMedia(media, collection string, factory func() Resource) {
	resources[Media(media)] = factory
	collections[Media(collection)] = factory
}

// Factory returns a resource of the specified media type
func Factory(media string) Resource {
	return resources[Media(media)]()
}
