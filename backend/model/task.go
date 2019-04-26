package model

// Task represents a task.
type Task struct {
	TaskID   string `db:"task_id" gqlgen:"id"`
	StatusID string `db:"status_id"`
	Title    string
	Position int
}
