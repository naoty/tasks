package model

// Status is the status of tasks.
type Status struct {
	StatusID string `gqlgen:"id"`
	Name     string
	Position int
	Tasks    []Task
}
