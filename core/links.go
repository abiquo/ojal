package core

// Links represents a list of links
type Links []*Link

// Rel returns the rel object link
func (l Links) Rel(rel string) (link *Link) {
	for _, l := range l {
		if l.Rel == rel {
			return l
		}
	}
	return
}

type FilterLink func(l *Link) bool

// Filter returns the links that meet the condition
func (l Links) Filter(filter FilterLink) (links Links) {
	for _, link := range l {
		if filter(link) {
			links = append(links, link)
		}
	}
	return
}

// Link retuns the self/edit link to thecore.Resource
func (l Links) Link() (link *Link) {
	if link = l.Rel("edit"); link == nil {
		link = l.Rel("self")
	}
	return
}

// URL returns a links collection href
func (l Links) URL() (u string) {
	if link := l.Link(); link != nil {
		u = link.Href
	}
	return
}

// Media returns a links collection media
func (l Links) Media() (m string) {
	if link := l.Link(); link != nil {
		m = link.Type
	}
	return
}

// DTO represents a list of links
type DTO struct {
	Links `json:"links,omitempty"`
}

// Add adds the Link l as rel to the *DTO Links
// If the link is nil, the *DTO Links do not change
func (d *DTO) Add(l *Link) {
	if l != nil {
		d.Links = append(d.Links, l)
	}
}

// NewDTO returns a DTO Links
func NewDTO(links ...*Link) (d DTO) {
	for _, link := range links {
		d.Add(link)
	}
	return
}

// Walk returns the DTO rel Resource
func (d *DTO) Walk(rel string) Resource {
	return d.Rel(rel).Walk()
}
