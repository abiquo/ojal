package core

import (
	"encoding/json"
	"net/url"
)

type page struct {
	Collection []json.RawMessage `json:"collection"`
	Size       int               `json:"totalSize"`
	DTO
}

// Collection represents an Abiquo API collection.
type Collection struct {
	next func() bool
	item func() Resource
	size func() int
}

// Next returns true it there are still elements on the collection
func (c *Collection) Next() bool { return c.next() }

// Item returns the first item and removes it from the collection
func (c *Collection) Item() Resource { return c.item() }

// Size returns the collection totalSize
func (c *Collection) Size() int { return c.size() }

// NewCollection returns a collection of the specified *Link.Type
func newCollection(link *Link, query url.Values) (collection *Collection) {
	next := link
	page := new(page)
	media := collections[link.Media()]
	factory := func(raw json.RawMessage) (resource Resource) {
		resource = Factory(media)
		json.Unmarshal(raw, resource)
		return
	}

	collection = &Collection{
		item: func() (resource Resource) {
			resource = factory(page.Collection[0])
			page.Collection = page.Collection[1:]
			return
		},
		next: func() bool {
			if len(page.Collection) == 0 && next != nil {
				Rest(page, Get(next, link).Query(query))
				next = page.Rel("next")
			}
			return len(page.Collection) > 0
		},
		size: func() int {
			return page.Size
		},
	}
	collection.Next()
	return
}

// List returns an slice with all the collection elements
func (c *Collection) List() (resources Resources) {
	for c.Next() {
		resources = append(resources, c.Item())
	}
	return
}

// Find a resource in a collection
func (c *Collection) Find(t func(r Resource) bool) Resource {
	for c.Next() {
		if resource := c.Item(); t(resource) {
			return resource
		}
	}
	return nil
}

// First returns a collection first element
func (c *Collection) First() Resource {
	return c.Find(func(Resource) bool { return true })
}
