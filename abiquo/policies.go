package abiquo

import "github.com/abiquo/ojal/core"

// FitPolicy represents a fit policy resource
type FitPolicy struct {
	FitPolicy string `json:"fitPolicy"` // Possible values are PROGRESSIVE and PERFORMANCE
	core.DTO
}

// MachineLoadRule represents a machine load rule resource
type MachineLoadRule struct {
	Aggregated        bool `json:"aggregated"`
	CPULoadPercentage int  `json:"cpuLoadPercentage"`
	RAMLoadPercentage int  `json:"ramLoadPercentage"`
	core.DTO
}

// DatastoreLoadRule represents a datastore load rule resource
type DatastoreLoadRule struct {
	StorageLoadPercentage int `json:"storageLoadPercentage"`
	core.DTO
}

// Rules represents a rules resource
type Rules struct {
	DatastoreLoadRules struct {
		Collection []*DatastoreLoadRule `json:"collection"`
		core.DTO
	} `json:"datastoreLoadRules"`
	EnterpriseExclusionRules struct {
		Collection []interface{} `json:"collection"`
		core.DTO
	} `json:"enterpriseExclusionRules"`
	FitPolicyRules struct {
		Collection []*FitPolicy `json:"collection"`
		core.DTO
	} `json:"fitPolicyrules"`
	MachineLoadRules struct {
		Collection []*MachineLoadRule `json:"collection"`
		core.DTO
	} `json:"machineLoadRules"`
	core.DTO
}
