package core

import (
	"errors"
	"net/url"
	"strings"
)

// Link represents an Abiquo link
type Link struct {
	Href  string `json:"href,omitempty"`
	Rel   string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`

	// Disk specific attributes
	Length             int    `json:"length,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	DiskController     string `json:"diskController,omitempty"`
	DiskLabel          string `json:"diskLabel,omitempty"`
}

// NewLink returns a new link to the href URL
func NewLink(href string) *Link {
	return &Link{Href: Resolve(href, nil)}
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
	if l != nil {
		return strings.HasPrefix(l.Type, getMedia(media))
	}
	return false
}

// SetType retuns a link clone with the specified type value
func (l *Link) SetType(media string) (link *Link) {
	if link = l.copy(); link != nil {
		link.Type = getMedia(media)
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
	return newCollection(l, query)
}

// Walk returns the resource the link is pointing to
func (l *Link) Walk() (resource Resource, err error) {
	if l == nil {
		return nil, errors.New("Walk: link is nil")
	}

	resource = resources[l.Type]()
	err = l.Read(resource)
	return

}

// Read ...
func (l *Link) Read(result interface{}) (err error) {
	_, err = Rest(result, Get(l, l))
	return
}

// Create ...
func (l *Link) Create(result interface{}) (err error) {
	_, err = Rest(result, Post(l, l, l, result))
	return
}

// Update ...
func (l *Link) Update(result interface{}) (err error) {
	_, err = Rest(result, Put(l, l, l, result))
	return
}

// Remove ...
func (l *Link) Remove() (err error) {
	_, err = Rest(nil, Delete(l))
	return
}

// Exists ...
func (l *Link) Exists() (bool, error) {
	err := l.Read(nil)
	return err == nil, err
}
