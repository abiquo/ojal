package abiquo

import "github.com/abiquo/ojal/core"

type FitPolicy struct {
	FitPolicy string `json:"fitPolicy"`
	core.DTO
}

type MachineLoadRule struct {
	Aggregated        bool `json:"aggregated"`
	CPULoadPercentage int  `json:"cpuLoadPercentage"`
	RAMLoadPercentage int  `json:"ramLoadPercentage"`
	core.DTO
}

type DatastoreLoadRule struct {
	StorageLoadPercentage int `json:"storageLoadPercentage"`
	core.DTO
}

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
