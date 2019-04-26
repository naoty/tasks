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
	initialStatus := model.Status{}
	err := r.DB.Get(&initialStatus, "SELECT * FROM statuses ORDER BY position ASC LIMIT 1")
	if err != nil {
		return nil, err
	}

	_, err = r.DB.NamedExec(`
		INSERT INTO
			tasks (status_id, title, position)
		SELECT
			:status_id
			, :title
			, (CASE
				WHEN MAX(position) IS NULL THEN 1
				ELSE MAX(position) + 1
			END)
		FROM
			tasks
		WHERE
			status_id = :status_id
	`, map[string]interface{}{
		"status_id": initialStatus.StatusID,
		"title":     input.Title,
	})
	if err != nil {
		return nil, err
	}

	task := model.Task{}
	err = r.DB.Get(&task, "SELECT * FROM tasks ORDER BY task_id DESC LIMIT 1")
	return &gqlgen.CreateTaskPayload{ClientMutationID: input.ClientMutationID, Task: task}, err
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input gqlgen.DeleteTaskInput) (*gqlgen.DeleteTaskPayload, error) {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE task_id = ?", input.ID)
	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteTaskPayload{ClientMutationID: input.ClientMutationID}, nil
}

func (r *mutationResolver) MoveTask(ctx context.Context, input gqlgen.MoveTaskInput) (*gqlgen.MoveTaskPayload, error) {
	return nil, nil
}
