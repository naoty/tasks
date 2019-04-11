package model

// Status is the status of tasks.
type Status struct {
	StatusID string
	Name     string
	Position int
	Tasks    []Task
}
