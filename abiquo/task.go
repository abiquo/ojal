package abiquo

import (
	"fmt"
	"time"

	"github.com/abiquo/ojal/core"
)

// Task represents an API task resource
type Task struct {
	UserID string `json:"userId"`
	TaskID string `json:"taskId"`
	State  string `json:"state"`
	core.DTO
}

// NewTask creates a new API task from the call
func NewTask(call *core.Call) (err error) {
	reply := new(core.DTO)
	if _, err = core.Rest(reply, call); err == nil {
		endpoint := reply.Rel("status").SetType("taskextended")
		if result := taskWait(endpoint); result != "FINISHED_SUCCESSFULLY" {
			err = fmt.Errorf("task: %v %v", endpoint.Href, result)
		}
	}
	return
}

func taskWait(endpoint *core.Link) string {
	task := new(Task)
	for {
		time.Sleep(10000 * time.Millisecond)
		core.Read(endpoint, task)
		switch task.State {
		case "FINISHED_SUCCESSFULLY":
			return task.State
		case "FINISHED_UNSUCCESSFULLY", "ABORTED", "CANCELLED", "ACK_ERROR":
			return task.State
		}
	}
}
