package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type queryResolver struct {
	*Root
}

func (r *queryResolver) Statuses(ctx context.Context) ([]model.Status, error) {
	statuses := []model.Status{}
	err := r.DB.Select(&statuses, "SELECT * FROM statuses ORDER BY position ASC")
	return statuses, err
}
