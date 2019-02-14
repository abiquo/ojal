package core

// Href is an string representing an Abiquo resource
type Href interface {
	URL() string
}

// Type ...
type Type interface {
	Media() string
}

// Endpoint represents and Abiquo APIcore.Resource (urls, links, objects or responses)
type Endpoint interface {
	Href
	Type
}

// Resource is an abstract interface for DTO objects
type Resource interface {
	Endpoint
	Link() *Link
	Rel(string) *Link
	Add(*Link)
}
