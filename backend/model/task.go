package model

// Task represents a task.
type Task struct {
	TaskID   string `gqlgen:"id"`
	Title    string
	StatusID string
}
