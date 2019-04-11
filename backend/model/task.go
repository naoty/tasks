package model

// Task represents a task.
type Task struct {
	TaskID   string `json:"id"`
	Title    string `json:"title"`
	StatusID string `json:"statusId"`
}
