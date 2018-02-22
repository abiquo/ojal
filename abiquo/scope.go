package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// ScopeEntity represents an Abiquo API scope entity DTO
type ScopeEntity struct {
	ID         int    `json:"id,omitempty"`
	IDResource int    `json:"idResource"`
	EntityType string `json:"type"`
	core.DTO
}

// Scope represents an Abiquo API scopecore.Resource
type Scope struct {
	AutomaticAddDatacenter bool          `json:"automaticAddDatacenter,omitempty"`
	AutomaticAddEnterprise bool          `json:"automaticAddEnterprise,omitempty"`
	Entities               []ScopeEntity `json:"scopeEntities"`
	ID                     int           `json:"id,omitempty"`
	Name                   string        `json:"name"`
	core.DTO
}

func (s *Scope) Create() error {
	return core.Create(core.NewLinker("admin/scopes", "scope"), s)
}

// Scopes returns the API scopes collection
func Scopes(query url.Values) *core.Collection {
	return core.NewLinker("admin/scopes", "scopes").Collection(query)
}
