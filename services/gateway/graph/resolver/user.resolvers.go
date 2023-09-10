package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/yuorei/hackathon/graph/generated"
	"github.com/yuorei/hackathon/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserPayload, error) {
	response, err := r.client.CreateUser(input)
	if err != nil {
		return nil, err
	}

	return &model.UserPayload{
		ID:    "response.Idです",
		Name:  response.Name,
		Email: response.Email,
		// Gender:
		// Affiliation: response.Affiliation,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.UserPayload, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
