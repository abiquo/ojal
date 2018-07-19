package abiquo

import "github.com/abiquo/ojal/core"

// Alert represents an Abiquo alert DTO
type Alert struct {
	Active      bool       `json:"active"`
	Description string     `json:"description"`
	Muted       bool       `json:"muted"`
	Name        string     `json:"name"`
	Subscribers []string   `json:"subscribers,omitempty"`
	Alarms      core.Links `json:"alarms"`
	core.DTO
}

// ScalingGroup represents an Abiquo API Scaling group DTO
type ScalingGroup struct {
	Name        string             `json:"name"`
	Cooldown    int                `json:"defaultCooldownSeconds"`
	Max         int                `json:"maxSize"`
	Min         int                `json:"minSize"`
	Maintenance bool               `json:"maintenanceMode"`
	ScaleIn     []ScalingGroupRule `json:"scaleInRules"`
	ScaleOut    []ScalingGroupRule `json:"scaleOutRules"`
	core.DTO
}

// ScalingGroupRule represents an Abiquo API scaling group rule DTO
type ScalingGroupRule struct {
	NumberOfInstances int   `json:"numberOfInstances"`
	StartTime         int64 `json:"startTime,omitempty"`
	EndTime           int64 `json:"endTime,omitempty"`
	core.DTO
}

// StartMaintenance enables the *ScalingGroup maintenance mode
func (s *ScalingGroup) StartMaintenance() (err error) {
	_, err = core.Rest(nil, core.Post(
		s.Rel("startmaintenance").Href,
		"application/json, text/plain, */*",
		"application/json",
		nil,
	))
	return
}

// EndMaintenance disables the *ScalingGroup maintenance mode
func (s *ScalingGroup) EndMaintenance() (err error) {
	_, err = core.Rest(nil, core.Post(
		s.Rel("endmaintenance").Href,
		"application/json, text/plain, */*",
		"application/json",
		nil,
	))
	return
}

// VirtualAppliance represents an Abiquo virtual appliance DTO
type VirtualAppliance struct {
	Name string `json:"name"`
	core.DTO
}
