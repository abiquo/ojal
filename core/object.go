package core

// Object represents a generic Abiquo API object
type Object map[string]interface{}

// Links returns an Abiquo API object Links array
func (o Object) links() (links Links) {
	if _, ok := o["links"]; ok {
		for _, link := range o["links"].([]interface{}) {
			attr := link.(map[string]interface{})
			links = append(links, &Link{
				Href:  attr["href"].(string),
				Rel:   attr["rel"].(string),
				Title: attr["title"].(string),
				Type:  attr["type"].(string),
			})
		}
	}
	return
}

// Rel returns the rel object link
func (o Object) Rel(rel string) (link *Link) { return o.links().Rel(rel) }

// Media returns an Object type
func (o Object) Media() (media string) { return o.links().Media() }

// Href returns an object Caller
func (o Object) Href() (url string) { return o.links().URL() }

// Add adds a link to the object
func (o Object) Add(link *Link) {
	if link != nil {
		if o["links"] == nil {
			o["links"] = []interface{}{}
		}
		links := o["links"].([]interface{})
		o["links"] = append(links, map[string]interface{}{
			"href":  link.Href,
			"rel":   link.Rel,
			"title": link.Title,
			"type":  link.Type,
		})
	}
}

// Link retuns the self/edit link for the Object
func (o Object) Link() (link *Link) {
	if link = o.Rel("edit"); link == nil {
		link = o.Rel("self")
	}
	return
}
