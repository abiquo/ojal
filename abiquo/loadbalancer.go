package abiquo

import "github.com/abiquo/opal/core"

type LoadBalancerAddress struct {
	Internal bool `json:"internal"`
}

type LoadBalancerAddresses struct {
	Collection []LoadBalancerAddress `json:"collection"`
}

type RoutingRule struct {
	PortIn      string `json:"portIn"`
	PortOut     string `json:"portOut"`
	ProtocolIn  string `json:"protocolIn"`
	ProtocolOut string `json:"protocolOut"`
}

type RoutingRules struct {
	Collection []RoutingRule `json:"collection"`
}

type LoadBalancer struct {
	Algorithm             string                `json:"algorithm"`
	HealthChecks          []interface{}         `json:"healthChecks,omitempty"`
	Name                  string                `json:"name"`
	LoadBalancerAddresses LoadBalancerAddresses `json:"loadBalancerAddresses"`
	RoutingRules          RoutingRules          `json:"routingRules"`
	core.DTO
}

func (l *LoadBalancer) SetRules(rules []RoutingRule) error {
	return core.Update(l.Rel("rules"), &RoutingRules{rules})
}
