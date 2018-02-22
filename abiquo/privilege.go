package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// Privilege represents an Abiquo API Privilege DTO
type Privilege struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Events []string `json:"events"`
	core.DTO
}

// Privileges retuns the API privileges collection
func Privileges(query url.Values) *core.Collection {
	return core.NewLinker("config/privileges", "privileges").Collection(query)
}
