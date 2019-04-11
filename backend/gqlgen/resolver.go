package gqlgen

import (
	"context"
	"database/sql"

	"github.com/naoty/tasks/backend/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	*sql.DB
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Task() TaskResolver {
	return &taskResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Statuses(ctx context.Context) ([]model.Status, error) {
	rows, err := r.DB.Query("SELECT * FROM statuses LEFT OUTER JOIN tasks USING (status_id) ORDER BY position ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type result struct {
		statusID string
		name     string
		position int
		taskID   sql.NullString
		title    sql.NullString
	}

	statusMap := map[string]model.Status{}

	for rows.Next() {
		var result result
		err := rows.Scan(&result.statusID, &result.name, &result.position, &result.taskID, &result.title)
		if err != nil {
			return nil, err
		}

		status, ok := statusMap[result.statusID]
		if !ok {
			status = model.Status{StatusID: result.statusID, Name: result.name, Position: result.position, Tasks: []model.Task{}}
		}

		if result.taskID.Valid {
			task := model.Task{TaskID: result.taskID.String, Title: result.title.String, StatusID: result.statusID}
			status.Tasks = append(status.Tasks, task)
		}

		statusMap[result.statusID] = status
	}

	statuses := []model.Status{}

	for _, status := range statusMap {
		statuses = append(statuses, status)
	}

	return statuses, nil
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (*model.Status, error) {
	panic("not implemented")
}
