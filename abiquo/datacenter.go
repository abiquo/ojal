package abiquo

import "github.com/abiquo/ojal/core"

// BackupConfiguration represents a backup configuration resource
type BackupConfiguration struct {
	Subtype string   `json:"subtype"`
	Time    string   `json:"time"`
	Type    string   `json:"type"`
	Days    []string `json:"days,omitempty"`
	Sources []string `json:"sources,omitempty"`
}

// BackupManager represents a datacenter backup manager resource
type BackupManager struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// BackupPluginType represents a backup plugin type resource
type BackupPluginType struct {
	Constraints map[string]string        `json:"constraints"`
	Name        string                   `json:"name"`
	Operations  map[string][]interface{} `json:"operations"`
	Type        string                   `json:"type"`
	core.DTO
}

// BackupPolicy represents an Abiquo backup policy DTO
type BackupPolicy struct {
	Code           string                `json:"code"`
	Configurations []BackupConfiguration `json:"backupConfigurations"`
	Name           string                `json:"name"`
	core.DTO
}

// BackupType represents an Abiquo backup type DTO
type BackupType struct {
	BackupTypes []map[string]string `json:"backupTypes"`
	core.DTO
}

// Datacenter represents an Abiquo datacenter DTO
//
// Collections:
// - backuppolicies
// - networks
// - hardwareprofiles
// - discover
type Datacenter struct {
	ID             int    `json:"id,omitempty"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	RemoteServices struct {
		Collection []RemoteService `json:"collection"`
		core.DTO
	} `json:"remoteServices"`
	core.DTO
}

// Device represents an Abiquo network SDN device DTO
type Device struct {
	Default     bool   `json:"vdcDefault"`
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
	Name        string `json:"name"`
	Password    string `json:"password,omitempty"`
	Username    string `json:"user,omitempty"`
	core.DTO
}

// HardwareProfile represents an Abiquo hardware profile DTO
type HardwareProfile struct {
	Active  bool   `json:"active"`
	Name    string `json:"name"`
	CPU     int    `json:"cpu"`
	RAMInMB int    `json:"ramInMb"`
	core.DTO
}

// Machine represents an Abiquo physical machine DTO
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
	Port     int    `json:"port,omitempty"`
	RAM      int    `json:"ram"`
	RAMUsed  int    `json:"ramUsed"`
	State    string `json:"state"`
	Type     string `json:"type"`
	User     string `json:"user"`
	core.DTO
}

// NetworkServiceType represents an Abiquo network service type DTO
type NetworkServiceType struct {
	Default bool   `json:"defaultNST"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	core.DTO
}

// Rack represents an Abiquo datacenter rack DTO
type Rack struct {
	Description string `json:"shortDescription"`
	HAEnabled   bool   `json:"haEnabled,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	NRSQ        int    `json:"nrsq,omitempty"`
	VlanIDMax   int    `json:"vlanIdMax,omitempty"`
	VlanIDMin   int    `json:"vlanIdMin,omitempty"`
	Reserved    int    `json:"vlanPerVdcReserved,omitempty"`
	core.DTO
}
