package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

type RemoteService struct {
	ID     int    `json:"id,omitempty"`
	Status int    `json:"status"`
	Type   string `json:"type"`
	URI    string `json:"uri"`
	UUID   string `json:"uuid"`
	core.DTO
}

func newRemoteService() core.Resource { return new(RemoteService) }

func RemoteServices(q url.Values) *core.Collection {
	return core.NewLinker("admin/remoteservices", "remoteservices").Collection(q)
}
