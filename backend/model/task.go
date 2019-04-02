package model

// Task represents a task.
type Task struct {
	TaskID   string `json:"taskId"`
	Title    string `json:"title"`
	StatusID string `json:"statusId"`
}
