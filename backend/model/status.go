package model

// Status is the status of tasks.
type Status struct {
	StatusID string   `json:"statusId"`
	Name     string   `json:"name"`
	TaskIDs  []string `json:"taskIds"`
}
