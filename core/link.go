package core

import (
	"errors"
	"net/url"
)

// Link represents an Abiquo link
type Link struct {
	Href  string `json:"href"`
	Rel   string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`

	DiskControllerType string `json:"diskControllerType,omitempty"`
	DiskController     string `json:"diskController,omitempty"`
}

// NewLink returns a new link to the href URL
func NewLink(href string) *Link {
	return &Link{Href: Resolve(href, nil)}
}

// NewLinkType returns a new link of the specified type
func NewLinkType(href, media string) *Link {
	return NewLink(href).SetType(media)
}

func (l *Link) copy() (link *Link) {
	if l != nil {
		link = new(Link)
		*link = *l
	}
	return
}

// URL returns a link href
func (l *Link) URL() (href string) {
	if l != nil {
		href = l.Href
	}
	return
}

// Media returns a link type value
func (l *Link) Media() (media string) {
	if l != nil {
		media = l.Type
	}
	return
}

// IsMedia returns whether the link is of media Type
func (l *Link) IsMedia(media string) bool {
	return l.Type == Media(media)
}

// SetType retuns a link clone with the specified type value
func (l *Link) SetType(media string) (link *Link) {
	if link = l.copy(); link != nil {
		link.Type = Media(media)
	}
	return
}

// SetRel retuns a link clone with the specified rel value
func (l *Link) SetRel(rel string) (link *Link) {
	if link = l.copy(); link != nil {
		link.Rel = rel
	}
	return
}

// SetTitle retuns a link clone with the specified tytle value
func (l *Link) SetTitle(title string) (link *Link) {
	if link = l.copy(); link != nil {
		link.Title = title
	}
	return
}

// Collection returns a collection from the link
func (l *Link) Collection(query url.Values) *Collection {
	return NewCollection(l, query)
}

// Walk returns the resource the link is pointing to
func (l *Link) Walk() (resource Resource, err error) {
	if l == nil {
		return nil, errors.New("Walk: link is nil")
	}

	r := resources[l.Type]()
	err = Read(l, r)
	if err != nil {
		return
	}

	resource = r
	return
}

// Exists returns whether a link is a valid resource in the API
func (l *Link) Exists() (exists bool, err error) {
	if l != nil {
		var resource interface{}
		err = Read(l, resource)
		exists = err == nil
	}
	return
}

// ToHref ...
func (l *Link) ToHref() string {
	return l.URL()
}
