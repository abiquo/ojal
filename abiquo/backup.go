package abiquo

import "github.com/abiquo/ojal/core"

type BackupManager struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type BackupConfiguration struct {
	Subtype string   `json:"subtype"`
	Time    string   `json:"time"`
	Type    string   `json:"type"`
	Days    []string `json:"days,omitempty"`
	Sources []string `json:"sources,omitempty"`
}

type BackupPluginType struct {
	Constraints map[string]string        `json:"constraints"`
	Name        string                   `json:"name"`
	Operations  map[string][]interface{} `json:"operations"`
	Type        string                   `json:"type"`
	core.DTO
}

type BackupPolicy struct {
	Code           string                `json:"code"`
	Configurations []BackupConfiguration `json:"backupConfigurations"`
	Name           string                `json:"name"`
	core.DTO
}

type BackupType struct {
	BackupTypes []map[string]string `json:"backupTypes"`
	core.DTO
}
