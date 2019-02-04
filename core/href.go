package core

// Href is an string representing an Abiquo resource
type Href string

// URL ...
func (h Href) URL() string {
	return string(h)
}

// Media ...
func (h Href) Media() string {
	return ""
}
