package abiquo

import "github.com/abiquo/opal/core"

type MachineInterface struct {
	MAC  string `json:"mac"`
	Name string `json:"name"`
	core.DTO
}

type Machine struct {
	AgentUser  string `json:"agentUser"`
	CPU        int    `json:"cpu"`
	CPUUsed    int    `json:"cpuUsed"`
	Datastores struct {
		Collection []Datastore `json:"collection"`
		core.DTO
	} `json:"datastores"`
	Description       string `json:"description"`
	Hostname          string `json:"hostname"`
	InitiatorIQN      string `json:"initiatorIQN"`
	IP                string `json:"ip"`
	IPService         string `json:"ipService"`
	ManagerPassword   string `json:"managerPassword"`
	ManagerUser       string `json:"managerUser"`
	Name              string `json:"name"`
	NetworkInterfaces struct {
		Collection []MachineInterface `json:"collection"`
		core.DTO
	} `json:"networkInterfaces"`
	Password string `json:"password"`
	RAM      int    `json:"ram"`
	RAMUsed  int    `json:"ramUsed"`
	State    string `json:"state"`
	Type     string `json:"type"`
	User     string `json:"user"`
	core.DTO
}

func newMachine() core.Resource { return new(Machine) }
