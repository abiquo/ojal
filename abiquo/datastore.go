package abiquo

import "github.com/abiquo/opal/core"

type Datastore struct {
	DatastoreUUID string `json:"datastoreUUID"`
	Directory     string `json:"directory"`
	Enabled       bool   `json:"enabled"`
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	RootPath      string `json:"rootPath"`
	Size          int64  `json:"size"`
	UsedSize      int64  `json:"usedSize"`
	core.DTO
}

func newDatastore() core.Resource { return new(Datastore) }

type DatastoreTier struct {
	DefaultAllowed          bool   `json:"defaultAllowed"`
	DefaultForDatacenter    bool   `json:"defaultForDatacenter"`
	Description             string `json:"description"`
	Enabled                 bool   `json:"enabled"`
	Name                    string `json:"name"`
	StorageAllocationPolicy string `json:"storageAllocationPolicy"`
	core.DTO
}

func newDatastoreTier() core.Resource { return new(DatastoreTier) }