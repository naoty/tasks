package resolver

import (
	"github.com/jmoiron/sqlx"
	"github.com/naoty/tasks/backend/gqlgen"
)

// Root is an implementation of gqlgen.ResolverRoot.
type Root struct {
	*sqlx.DB
}

// Mutation returns a resolver for mutation.
func (r *Root) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}

// Query returns a resolver for query.
func (r *Root) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

// Status returns a resolver for statuses.
func (r *Root) Status() gqlgen.StatusResolver {
	return &statusResolver{r}
}

// Task returns a resolver for tasks.
func (r *Root) Task() gqlgen.TaskResolver {
	return &taskResolver{r}
}
