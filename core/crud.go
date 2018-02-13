package core

// Create posts the resource to the endpoint
func Create(endpoint Endpoint, resource interface{}) error {
	_, e := Rest(resource, Post(
		endpoint.URL(),
		endpoint.Media(),
		endpoint.Media(),
		resource,
	))
	return e
}

// Read gets the resource from the endpoint
func Read(endpoint Endpoint, resource interface{}) error {
	_, e := Rest(resource, Get(
		endpoint.URL(),
		endpoint.Media(),
	))
	return e
}

// Update puts the resource into the endpoint
func Update(endpoint Endpoint, resource interface{}) error {
	_, e := Rest(resource, Put(
		endpoint.URL(),
		endpoint.Media(),
		endpoint.Media(),
		resource,
	))
	return e
}

// Remove deletes de resource in the endpoint
func Remove(endpoint Endpoint) error {
	_, e := Rest(nil, Delete(
		endpoint.URL(),
	))
	return e
}
