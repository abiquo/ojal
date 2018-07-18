package abiquo

import "github.com/abiquo/ojal/core"

type ActionPlan struct {
	CreatedBy   string            `json:"createdBy,omitempty"`
	Description string            `json:"description"`
	Entries     []ActionPlanEntry `json:"entries"`
	Name        string            `json:"name"`
	core.DTO
}

type ActionPlanEntry struct {
	Parameter     string `json:"parameter"`
	ParameterType string `json:"parameterType"`
	Sequence      int    `json:"sequence"`
	Type          string `json:"type"`
}
