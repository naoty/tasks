package model

// Status is the status of tasks.
type Status struct {
	StatusID string `db:"status_id" gqlgen:"id"`
	Name     string
	Position int
}
