package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

type RemoteService struct {
	ID     int    `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Type   string `json:"type"`
	URI    string `json:"uri"`
	UUID   string `json:"uuid,omitempty"`
	core.DTO
}

func RemoteServices(q url.Values) *core.Collection {
	return core.NewLinker("admin/remoteservices", "remoteservices").Collection(q)
}
