package core

import (
	"bytes"
	"fmt"
)

// Error represents an Abiquo API Error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Links
}

// Errors represents an Abiquo API error collection
type Errors struct {
	status     int
	msg        string
	Collection []Error `json:"collection"`
	Links
}

func (e *Error) Error() string {
	return e.Code + " " + e.Message
}

func (e *Errors) Error() string {
	message := new(bytes.Buffer)
	fmt.Fprint(message, e.msg)
	for _, err := range e.Collection {
		fmt.Fprint(message, ": ", err.Error())
	}
	return string(message.Bytes())
}
