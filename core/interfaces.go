package core

import "net/url"

// Endpoint represents and Abiquo APIcore.Resource (urls, links, objects or responses)
type Endpoint interface {
	URL() string
	Media() string
}

// Linker represents links which can generate collections or being walked
type Linker interface {
	Endpoint
	Walk() (Resource, error)
	Collection(q url.Values) *Collection
}

// NewLinker returns an abstract endpoint. Useful for Walk and Collection
func NewLinker(h, m string) Linker {
	return NewLinkType(h, m)
}
