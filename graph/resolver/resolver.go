package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/MitsuguSueyoshi/graphql-golang-todos/graph/generated"
	"github.com/MitsuguSueyoshi/graphql-golang-todos/graph/model"
)

type Resolver struct{}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	name := "Bob"
	return &model.User{
		ID:   *id,
		Name: name,
	}, nil
}

func (r *queryResolver) Dog(ctx context.Context, id *string) (*model.Dog, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
