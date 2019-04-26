package model

// Task represents a task.
type Task struct {
	TaskID   string `gqlgen:"id"`
	StatusID string
	Title    string
	Position int
}
