package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang_server_practice/graph/generated"
	"golang_server_practice/graph/model"
)

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	dummyUser := &model.User{ID: "1", Name: "jeff"}

	var result []*model.User
	result = append(result, dummyUser)

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
