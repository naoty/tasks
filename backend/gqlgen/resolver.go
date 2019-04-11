package gqlgen

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Status() StatusResolver {
	return &statusResolver{r}
}
func (r *Resolver) Task() TaskResolver {
	return &taskResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Statuses(ctx context.Context) ([]model.Status, error) {
	statuses := []model.Status{
		model.Status{StatusID: "1", Name: "TODO", Position: 1, Tasks: []model.Task{
			model.Task{TaskID: "1", Title: "dummy1", StatusID: "1"},
			model.Task{TaskID: "2", Title: "dummy2", StatusID: "1"},
		}},
		model.Status{StatusID: "2", Name: "DOING", Position: 2, Tasks: []model.Task{
			model.Task{TaskID: "3", Title: "dummy3", StatusID: "2"},
		}},
		model.Status{StatusID: "3", Name: "DONE", Position: 3, Tasks: []model.Task{}},
	}

	return statuses, nil
}

type statusResolver struct{ *Resolver }

func (r *statusResolver) ID(ctx context.Context, obj *model.Status) (string, error) {
	return obj.StatusID, nil
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (string, error) {
	return obj.TaskID, nil
}
func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (*model.Status, error) {
	panic("not implemented")
}
