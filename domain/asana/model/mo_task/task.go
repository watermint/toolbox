package mo_task

import "encoding/json"

type Task struct {
	Raw          json.RawMessage
	Gid          string `json:"gid" path:"gid"`
	Name         string `json:"name" path:"name"`
	ResourceType string `json:"resource_type" path:"resource_type"`
	CreatedAt    string `json:"created_at" path:"created_at"`
	Completed    bool   `json:"completed" path:"completed"`
	CompletedAt  string `json:"completed_at" path:"completed_at"`
	DueAt        string `json:"due_at" path:"due_at"`
	DueOn        string `json:"due_on" path:"due_on"`
}
