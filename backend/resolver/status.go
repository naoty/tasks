package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type statusResolver struct {
	*Root
}

func (r *statusResolver) Tasks(ctx context.Context, obj *model.Status) ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.DB.Select(&tasks, "SELECT * FROM tasks WHERE status_id = ? ORDER BY task_id ASC", obj.StatusID)
	return tasks, err
}
