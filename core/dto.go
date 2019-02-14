package core

// DTO represents a list of links
type DTO struct {
	Links `json:"links,omitempty"`
}

// Add adds the Link l as rel to the *DTO Links
// If the link is nil, the *DTO Links do not change
func (d *DTO) Add(l *Link) {
	if l == nil {
		return
	}
	d.Links = append(d.Links, l)
}

// NewDTO returns a DTO Links
func NewDTO(links ...*Link) (d DTO) {
	for _, link := range links {
		d.Add(link)
	}
	return
}

// Read ...
func (d *DTO) Read(result interface{}) (err error) {
	_, err = Rest(result, Get(d, d))
	return
}

// Create ...
func (d *DTO) Create(result interface{}) (err error) {
	_, err = Rest(result, Post(d, d, d, result))
	return
}

// Update ...
func (d *DTO) Update(result interface{}) (err error) {
	_, err = Rest(result, Put(d, d, d, result))
	return
}

// Remove ...
func (d *DTO) Remove() (err error) {
	_, err = Rest(nil, Delete(d))
	return
}

// Exists ...
func (d *DTO) Exists() (bool, error) {
	err := d.Read(nil)
	return err == nil, err
}
