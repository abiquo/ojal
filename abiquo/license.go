package abiquo

import (
	"net/url"

	"github.com/abiquo/opal/core"
)

type License struct {
	Code                 string `json:"code"`
	Expiration           string `json:"expiration,omitempty"`
	ID                   int    `json:"id,omitempty"`
	NumCores             int    `json:"numcores,omitempty"`
	ScalingGroupsEnabled bool   `json:"scalingGroupsEnabled,omitempty"`
	core.DTO
}

func newLicense() core.Resource { return new(License) }

func Licenses(query url.Values) *core.Collection {
	return core.NewLinker("config/licenses", "licenses").Collection(query)
}
