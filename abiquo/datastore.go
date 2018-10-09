package abiquo

import "github.com/abiquo/ojal/core"

// Datastore represents a datastore resource
type Datastore struct {
	UUID      string `json:"datastoreUUID"`
	Directory string `json:"directory"`
	Enabled   bool   `json:"enabled"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	RootPath  string `json:"rootPath"`
	Size      int64  `json:"size"`
	UsedSize  int64  `json:"usedSize"`
	core.DTO
}

// DatastoreTier represents a datastore tier resource
type DatastoreTier struct {
	DefaultAllowed       bool   `json:"defaultAllowed"`
	DefaultForDatacenter bool   `json:"defaultForDatacenter"`
	Description          string `json:"description"`
	Enabled              bool   `json:"enabled"`
	Name                 string `json:"name"`
	Policy               string `json:"storageAllocationPolicy"`
	core.DTO
}
