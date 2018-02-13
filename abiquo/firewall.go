package abiquo

import . "github.com/abiquo/opal/core"

type Firewall struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProviderId  string `json:"providerId,omitempty"`
	DTO
}

type FirewallRule struct {
	Id         int      `json:"id,omitempty"`
	FromPort   int      `json:"fromPort"`
	ToPort     int      `json:"toPort"`
	Protocol   string   `json:"protocol"`
	ProviderId string   `json:"providerId"`
	Sources    []string `json:"sources"`
	Targets    []string `json:"targets"`
}

type FirewallRules struct {
	Collection []FirewallRule `json:"collection"`
}

func (f *Firewall) SetRules(rules *FirewallRules) error {
	return Update(f.Rel("rules"), rules)
}
