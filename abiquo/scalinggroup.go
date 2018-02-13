package abiquo

import (
	. "github.com/abiquo/opal/core"
)

// ScalingGroup represents an Abiquo API Scaling group DTO
type ScalingGroup struct {
	Name        string             `json:"name"`
	Cooldown    int                `json:"defaultCooldownSeconds"`
	Max         int                `json:"maxSize"`
	Min         int                `json:"minSize"`
	Maintenance bool               `json:"maintenanceMode"`
	ScaleIn     []ScalingGroupRule `json:"scaleInRules"`
	ScaleOut    []ScalingGroupRule `json:"scaleOutRules"`
	DTO
}

// ScalingGroupRule represents an Abiquo API scaling group rule DTO
type ScalingGroupRule struct {
	NumberOfInstances int   `json:"numberOfInstances"`
	StartTime         int64 `json:"startTime,omitempty"`
	EndTime           int64 `json:"endTime,omitempty"`
	DTO
}

// StartMaintenance enables the *ScalingGroup maintenance mode
func (s *ScalingGroup) StartMaintenance() (err error) {
	_, err = Rest(nil, Post(
		s.Rel("startmaintenance").Href,
		"application/json, text/plain, */*",
		"application/json",
		nil,
	))
	return
}
