package abiquo

import (
	"net/url"

	"github.com/abiquo/ojal/core"
)

// Device represents a network SDN device DTO
type Device struct {
	Default     bool   `json:"vdcDefault"`
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Username    string `json:"user"`
	core.DTO
}

type DeviceInterface struct {
	Interface   string            `json:"deviceInterface"`
	RealName    string            `json:"realName"`
	Constraints map[string]string `json:"constraints"`
	Operations  []interface{}     `json:"operations"`
}

type DeviceType struct {
	Name       string            `json:"name"`
	Interfaces []DeviceInterface `json:"deviceInterfaces"`
	core.DTO
}

// DeviceTypes returns the API supported SDN device types collection
func DeviceTypes(q url.Values) *core.Collection {
	return core.NewLinker("config/devicetypes", "devicetypes").Collection(q)
}
