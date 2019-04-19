package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/gqlgen"

	"github.com/naoty/tasks/backend/model"
)

type mutationResolver struct {
	*Root
}

func (r *mutationResolver) CreateTask(ctx context.Context, input gqlgen.CreateTaskInput) (*gqlgen.CreateTaskPayload, error) {
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

	_, err = r.DB.Exec("INSERT INTO tasks (title, status_id) VALUES (?, ?)", input.Title, initialStatusID)
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

	return &gqlgen.CreateTaskPayload{ClientMutationID: input.ClientMutationID, Task: task}, err
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input gqlgen.DeleteTaskInput) (*gqlgen.DeleteTaskPayload, error) {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE task_id = ?", input.ID)
	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteTaskPayload{ClientMutationID: input.ClientMutationID}, nil
}
