package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type mutationResolver struct {
	*Root
}

func (r *mutationResolver) CreateTask(ctx context.Context, title string) (*model.Task, error) {
	panic("not implemented")
}
