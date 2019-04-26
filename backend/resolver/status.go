package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type statusResolver struct {
	*Root
}

func (r *statusResolver) Tasks(ctx context.Context, obj *model.Status) ([]model.Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE status_id = ? ORDER BY task_id ASC", obj.StatusID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.TaskID, &task.StatusID, &task.Title, &task.Position)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
