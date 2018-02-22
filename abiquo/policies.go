package abiquo

import "github.com/abiquo/opal/core"

type FitPolicy struct {
	FitPolicy string `json:"fitPolicy"`
	core.DTO
}

func NewFitPolicy() core.Resource { return new(FitPolicy) }

type MachineLoadRule struct {
	Aggregated        bool `json:"aggregated"`
	CPULoadPercentage int  `json:"cpuLoadPercentage"`
	RAMLoadPercentage int  `json:"ramLoadPercentage"`
	core.DTO
}

func NewMachineLoadRule() core.Resource { return new(MachineLoadRule) }

type DatastoreLoadRule struct {
	StorageLoadPercentage int `json:"storageLoadPercentage"`
	core.DTO
}

func NewDatastoreLoadRule() core.Resource { return new(DatastoreLoadRule) }

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
