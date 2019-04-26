package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type taskResolver struct {
	*Root
}

func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (*model.Status, error) {
	status := model.Status{}
	err := r.DB.Get(&status, "SELECT * FROM statuses WHERE status_id = ? ORDER BY status_id ASC LIMIT 1", obj.StatusID)
	return &status, err
}
