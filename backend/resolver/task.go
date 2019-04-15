package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type taskResolver struct {
	*Root
}

func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (*model.Status, error) {
	rows, err := r.DB.Query("SELECT * FROM statuses WHERE status_id = ? ORDER BY status_id ASC LIMIT 1", obj.StatusID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var status model.Status
	rows.Next()
	err = rows.Scan(&status.StatusID, &status.Name, &status.Position)

	return &status, err
}
