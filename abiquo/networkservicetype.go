package abiquo

import "github.com/abiquo/ojal/core"

type NetworkServiceType struct {
	Default bool   `json:"defaultNST"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	core.DTO
}
