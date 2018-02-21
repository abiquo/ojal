package abiquo

import "github.com/abiquo/opal/core"

type Machine struct {
	AgentUser  string `json:"agentUser"`
	CPU        int    `json:"cpu"`
	CPUUsed    int    `json:"cpuUsed"`
	Datastores struct {
		Collection []*Datastore `json:"collection"`
		core.DTO
	} `json:"datastores"`
	Description string `json:"description"`
	Hostname    string `json:"hostname"`
	Initiator   string `json:"initiatorIQN"`
	IP          string `json:"ip"`
	IPService   string `json:"ipService"`
	ManagerIP   string `json:"managerIp,omitempty"`
	ManagerPass string `json:"managerPassword,omitempty"`
	ManagerUser string `json:"managerUser,omitempty"`
	Name        string `json:"name"`
	Interfaces  struct {
		Collection []*struct {
			MAC  string `json:"mac"`
			Name string `json:"name"`
			core.DTO
		} `json:"collection"`
		core.DTO
	} `json:"networkInterfaces"`
	Password string `json:"password,omitempty"`
	Port     string `json:"port,omitempty"`
	RAM      int    `json:"ram"`
	RAMUsed  int    `json:"ramUsed"`
	State    string `json:"state"`
	Type     string `json:"type"`
	User     string `json:"user"`
	core.DTO
}

func newMachine() core.Resource { return new(Machine) }
