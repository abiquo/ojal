package abiquo

import "github.com/abiquo/ojal/core"

type Limit struct {
	CPUSoft  int `json:"cpuSoft"`
	HDSoft   int `json:"HDSoft"`
	IPSoft   int `json:"ipSoft"`
	RAMSoft  int `json:"ramSoft"`
	RepoSoft int `json:"repositorySoftInMb"`
	VolSoft  int `json:"volSoft"`
	VLANSoft int `json:"vlansSoft"`

	CPUHard  int `json:"cpuHard"`
	HDHard   int `json:"HDHard"`
	IPHard   int `json:"ipHard"`
	RAMHard  int `json:"ramHard"`
	RepoHard int `json:"repositoryHardInMb"`
	VolHard  int `json:"volHard"`
	VLANHard int `json:"vlansHard"`
	core.DTO
}
