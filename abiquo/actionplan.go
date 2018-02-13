package abiquo

import . "github.com/abiquo/opal/core"

type ActionPlan struct {
	CreatedBy   string            `json:"createdBy,omitempty"`
	Description string            `json:"description"`
	Entries     []ActionPlanEntry `json:"entries"`
	Name        string            `json:"name"`
	DTO
}

type ActionPlanEntry struct {
	Parameter     string `json:"parameter"`
	ParameterType string `json:"parameterType"`
	Sequence      int    `json:"sequence"`
	Type          string `json:"type"`
}

// SetTriggers posts a list of lists to *a alerttriggers link
func (a *ActionPlan) SetTriggers(d *DTO) error {
	return Create(a.Rel("alerttriggers"), d)
}
