package abiquo

import (
	"fmt"
	"time"

	"github.com/abiquo/ojal/core"
)

// Task represents an API task resource
type Task struct {
	CreationTimestamp int64  `json:"creationTimestamp"`
	OwnerID           string `json:"ownerId"`
	State             string `json:"state"`
	TaskID            string `json:"taskId"`
	Timestamp         int64  `json:"timestamp"`
	Type              string `json:"type"`
	UserID            string `json:"userId"`

	Jobs struct {
		TotalSize  interface{} `json:"totalSize"`
		Collection []struct {
			ID                string `json:"id"`
			ParentTaskID      string `json:"parentTaskId"`
			Type              string `json:"type"`
			Description       string `json:"description"`
			State             string `json:"state"`
			RollbackState     string `json:"rollbackState"`
			CreationTimestamp int64  `json:"creationTimestamp"`
			Timestamp         int64  `json:"timestamp"`
		} `json:"collection"`
	} `json:"jobs"`

	core.DTO
}

// NewTask creates a new API task from the call
func NewTask(call *core.Call) (err error) {
	reply := new(core.DTO)
	_, err = core.Rest(reply, call)
	if err != nil {
		return
	}

	endpoint := reply.Rel("status").SetType("taskextended")
	result, err := taskWait(endpoint)
	if err != nil {
		return
	}

	if result != "FINISHED_SUCCESSFULLY" {
		err = fmt.Errorf("task: %v %v", endpoint.Href, result)
	}
	return
}

func taskWait(endpoint *core.Link) (state string, err error) {
	task := new(Task)
	for {
		time.Sleep(10000 * time.Millisecond)
		err = endpoint.Read(task)
		if err != nil {
			return
		}

		state = task.State
		switch state {
		case "FINISHED_SUCCESSFULLY":
			return
		case "FINISHED_UNSUCCESSFULLY", "ABORTED", "CANCELLED", "ACK_ERROR":
			return
		}
	}
}
