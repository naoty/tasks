package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type mutationResolver struct {
	*Root
}

func (r *mutationResolver) CreateTask(ctx context.Context, title string) (*model.Task, error) {
	statusRows, err := r.DB.Query("SELECT status_id FROM statuses ORDER BY position ASC LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statusRows.Close()

	var initialStatusID int
	statusRows.Next()
	err = statusRows.Scan(&initialStatusID)
	if err != nil {
		return nil, err
	}

	_, err = r.DB.Exec("INSERT INTO tasks (title, status_id) VALUES (?, ?)", title, initialStatusID)
	if err != nil {
		return nil, err
	}

	taskRows, err := r.DB.Query("SELECT * FROM tasks ORDER BY task_id DESC LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer taskRows.Close()

	var task model.Task
	taskRows.Next()
	err = taskRows.Scan(&task.TaskID, &task.Title, &task.StatusID)

	return &task, err
}
