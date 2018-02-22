package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// Category represents and Abiquo API Category DTO
type Category struct {
	Name     string `json:"name"`
	Default  bool   `json:"defaultCategory"`
	ID       int    `json:"id,omitempty"`
	Erasable bool   `json:"erasable"`
	core.DTO
}

// NewCategory category retuns an new Abiquo API Category DTO
func NewCategory() core.Resource { return new(Category) }

func Categories(query url.Values) *core.Collection {
	return core.NewLinker("config/categories", "categories").Collection(query)
}

// Create creates a new Category in the Abiquo API
func (c *Category) Create() error {
	return core.Create(core.NewLinker("config/categories", "category"), c)
}
