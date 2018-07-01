package abiquo

import "github.com/abiquo/ojal/core"

// Limit represents an abiquo enterprise limit
type Limit struct {
	CPUHard   int  `json:"cpuHard"`
	CPUSoft   int  `json:"cpuSoft"`
	HDHard    int  `json:"HDHard"`
	HDSoft    int  `json:"HDSoft"`
	EnableHPs bool `json:"enabledHardwareProfiles"`
	IPHard    int  `json:"ipHard"`
	IPSoft    int  `json:"ipSoft"`
	RAMHard   int  `json:"ramHard"`
	RAMSoft   int  `json:"ramSoft"`
	RepoHard  int  `json:"repositoryHardInMb"`
	RepoSoft  int  `json:"repositorySoftInMb"`
	VLANHard  int  `json:"vlansHard"`
	VLANSoft  int  `json:"vlansSoft"`
	VolHard   int  `json:"volHard"`
	VolSoft   int  `json:"volSoft"`
	core.DTO
}
