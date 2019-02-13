package core

import (
	"encoding/json"
	"fmt"
)

// Error represents an Abiquo API error collection
type Error struct {
	Status     int
	Collection []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		// DTO
	} `json:"collection"`
	// DTO
}

func (e Error) Error() string {
	return fmt.Sprint(e.Collection)
}

func newError(code int, body []byte) (e Error) {
	e.Status = code
	json.Unmarshal(body, &e)
	return
}
