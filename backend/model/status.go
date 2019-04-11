package model

// Status is the status of tasks.
type Status struct {
	StatusID string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	Tasks    []Task `json:"tasks"`
}
