package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

// Role represents an Abiquo Rolecore.Resource
type Role struct {
	Blocked       bool     `json:"blocked"`
	ExternalRoles []string `json:"externalRoles,omitempty"` // PENDING
	ID            int      `json:"id,omitempty"`            // The role id
	IDEnterprise  int      `json:"idEnterprise,omitempty"`  // The enterprise of the role
	Name          string   `json:"name"`                    // The role name
	core.DTO
}

func Roles(query url.Values) *core.Collection {
	return core.NewLinker("admin/roles", "roles").Collection(query)
}

// Create posts the *Role r to the Abiquo API roles endpoint
func (r *Role) Create() error {
	return core.Create(core.NewLinker("admin/roles", "role"), r)
}

func (r *Role) Privileges(query url.Values) *core.Collection {
	return r.Rel("privileges").Collection(query)
}

// AddPrivilege adds the *Privilege rel privilege link to the *Role
func (r *Role) AddPrivilege(p *Privilege) {
	r.Add(p.Rel("privilege"))
}
