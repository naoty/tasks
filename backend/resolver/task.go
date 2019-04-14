package resolver

import (
	"context"

	"github.com/naoty/tasks/backend/model"
)

type taskResolver struct {
	*Root
}

func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (*model.Status, error) {
	panic("not implemented")
}
