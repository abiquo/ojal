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

// Find a resource in a collection
func (r Resources) Find(t Test) Resource {
	for _, resource := range r {
		if t(resource) {
			return resource
		}
	}
	return nil
}

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
