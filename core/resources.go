package core

// Resource is an abstract interface for DTO objects
type Resource interface {
	Endpoint
	Link() *Link
	Rel(string) *Link
	Add(*Link)
	Walk(string) Resource
}

// Resources represents an Abiquo API collection elements
type Resources []Resource

// Filter returns the elements which fullfill the Test
func (r Resources) Filter(t Test) (resources Resources) {
	for _, resource := range r {
		if t(resource) {
			resources = append(resources, resource)
		}
	}
	return
}

// Map applies a function to all elements of r
func (r Resources) Map(f func(Resource)) {
	for _, resource := range r {
		f(resource)
	}
}
