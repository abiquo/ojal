package abiquo

import "github.com/abiquo/ojal/core"

// RemoteService represents an Abiquo Remote Service DTO
type RemoteService struct {
	ID     int    `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Type   string `json:"type"`
	URI    string `json:"uri"`
	UUID   string `json:"uuid,omitempty"`
	core.DTO
}

// Role represents an Abiquo Role DTO
//
// Collections:
// - privileges
type Role struct {
	Blocked       bool     `json:"blocked"`
	ExternalRoles []string `json:"externalRoles,omitempty"` // PENDING
	ID            int      `json:"id,omitempty"`            // The role id
	IDEnterprise  int      `json:"idEnterprise,omitempty"`  // The enterprise of the role
	Name          string   `json:"name"`                    // The role name
	core.DTO
}

// Scope represents an Abiquo API scope DTO
type Scope struct {
	AutomaticAddDatacenter bool          `json:"automaticAddDatacenter,omitempty"`
	AutomaticAddEnterprise bool          `json:"automaticAddEnterprise,omitempty"`
	Entities               []ScopeEntity `json:"scopeEntities"`
	ID                     int           `json:"id,omitempty"`
	Name                   string        `json:"name"`
	core.DTO
}

// ScopeEntity represents an Abiquo API scope entity DTO
type ScopeEntity struct {
	ID         int    `json:"id,omitempty"`
	IDResource int    `json:"idResource"`
	EntityType string `json:"type"`
	core.DTO
}
