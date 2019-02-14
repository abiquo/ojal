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
	} `json:"collection"`
}

func (e Error) Error() (str string) {
	switch len(e.Collection) {
	case 0:
		return fmt.Sprintf("%v Unexpected status code", e.Status)
	case 1:
		return fmt.Sprintf("%v %v %v", e.Status, e.Collection[0].Code, e.Collection[0].Message)
	default:
		return fmt.Sprintf("%v Unexpected status code: %v", e.Status, e.Collection)
	}
}

func newError(code int, body []byte) (e Error) {
	e.Status = code
	json.Unmarshal(body, &e)
	return
}
