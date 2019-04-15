package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type queryResolver struct {
	*Root
}

func (r *queryResolver) Statuses(ctx context.Context) ([]model.Status, error) {
	rows, err := r.DB.Query("SELECT * FROM statuses ORDER BY position ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	statuses := []model.Status{}
	for rows.Next() {
		var status model.Status
		err := rows.Scan(&status.StatusID, &status.Name, &status.Position)
		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}
