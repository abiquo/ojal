package core

// Links represents a list of links
type Links []*Link

// Rel returns the rel object link
func (l Links) Rel(rel string) (link *Link) {
	return l.Find(func(l *Link) bool {
		return l.Rel == rel
	})
}

// Filter returns the links that meet the condition
func (l Links) Filter(t func(l *Link) bool) (links Links) {
	for _, link := range l {
		if t(link) {
			links = append(links, link)
		}
	}
	return
}

// Find returns the links that meet the condition
func (l Links) Find(t func(l *Link) bool) (link *Link) {
	for _, link := range l {
		if t(link) {
			return link
		}
	}
	return
}

// Map executes function f(*Link) on all the Links elements
func (l Links) Map(f func(*Link)) {
	for _, link := range l {
		f(link)
	}
}

// Link retuns the self/edit link to thecore.Resource
func (l Links) Link() (link *Link) {
	if link = l.Rel("edit"); link == nil {
		link = l.Rel("self")
	}
	return
}

// URL returns a links collection href
func (l Links) URL() (u string) { return l.Link().URL() }

// Media returns a links collection media
func (l Links) Media() (m string) { return l.Link().Media() }
