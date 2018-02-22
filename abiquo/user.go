package abiquo

import "github.com/abiquo/opal/core"

type User struct {
	Active      bool   `json:"active,omitempty"`
	AuthType    string `json:"authType,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email"`
	FirstLogin  bool   `json:"firstLogin,omitempty"`
	Locale      string `json:"locale"`
	Locked      bool   `json:"locked,omitempty"`
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick,omitempty"`
	Password    string `json:"password,omitempty"`
	Surname     string `json:"surname,omitempty"`
	core.DTO
}

func (u *User) Enterprise() (enterprise *Enterprise) {
	if e := u.Walk("enterprise"); e != nil {
		enterprise = e.(*Enterprise)
	}
	return
}
