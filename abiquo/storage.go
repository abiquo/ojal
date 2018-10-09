package abiquo

import "github.com/abiquo/ojal/core"

// StorageDevice represents a storage device resource
type StorageDevice struct {
	ID             int64  `json:"id,omitempty"`
	Name           string `json:"name"`
	Technology     string `json:"storageTechnology"`
	ManagementIP   string `json:"managementIp"`
	ManagementPort int64  `json:"managementPort"`
	ServiceIP      string `json:"serviceIp"`
	ServicePort    int64  `json:"servicePort"`
	core.DTO
}

// StoragePool represents a storage pool resource
type StoragePool struct {
	IDStorage         string `json:"idStorage,omitempty"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	TotalSizeInMB     int64  `json:"totalSizeInMb"`
	UsedSizeInMB      int64  `json:"usedSizeInMb"`
	AvailableSizeInMB int64  `json:"availableSizeInMb"`
	Enabled           bool   `json:"enabled"`
	MaxVolumes        int64  `json:"maxVolumes"`
	UsablePercent     int64  `json:"usablePercent"`
	core.DTO
}
