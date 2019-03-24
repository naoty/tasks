package model

// Status is the status of tasks.
type Status struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	TaskIDs []string `json:"taskIds"`
}
