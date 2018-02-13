package abiquo

import (
	"fmt"
	"time"

	. "github.com/abiquo/opal/core"
)

type Task struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
	State  string `json:"state"`
	DTO
}

// NewTask creates a new Abiquo API task from the caller
func NewTask(link *Link, media string, data interface{}) (err error) {
	reply := new(DTO)
	if _, err = Rest(reply, Post(
		link.Href,
		"acceptedrequest",
		media,
		data,
	)); err == nil {
		endpoint := reply.Rel("status").SetType("taskextended")
		if result := taskWait(endpoint); result != "FINISHED_SUCCESSFULLY" {
			err = fmt.Errorf("task: %v %v", endpoint.Href, result)
		}
	}
	return
}

func taskWait(endpoint *Link) string {
	task := new(Task)
	for {
		time.Sleep(10000 * time.Millisecond)
		Read(endpoint, task)
		switch task.State {
		case "FINISHED_SUCCESSFULLY":
			return task.State
		case "FINISHED_UNSUCCESSFULLY", "ABORTED", "CANCELLED", "ACK_ERROR":
			return task.State
		}
	}
}
