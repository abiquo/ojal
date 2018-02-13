package abiquo

import (
	"net/url"

	. "github.com/abiquo/opal/core"
)

// ScopeEntity represents an Abiquo API scope entity DTO
type ScopeEntity struct {
	ID         int    `json:"id,omitempty"`
	IDResource int    `json:"idResource"`
	EntityType string `json:"type"`
	DTO
}

// Scope represents an Abiquo API scopecore.Resource
type Scope struct {
	AutomaticAddDatacenter bool          `json:"automaticAddDatacenter,omitempty"`
	AutomaticAddEnterprise bool          `json:"automaticAddEnterprise,omitempty"`
	Entities               []ScopeEntity `json:"scopeEntities"`
	ID                     int           `json:"id,omitempty"`
	Name                   string        `json:"name"`
	DTO
}

func newScope() Resource { return new(Scope) }

func (s *Scope) Create() error {
	return Create(NewLinker("admin/scopes", "scope"), s)
}

// Scopes returns the API scopes collection
func Scopes(query url.Values) *Collection {
	return NewLinker("admin/scopes", "scopes").Collection(query)
}
