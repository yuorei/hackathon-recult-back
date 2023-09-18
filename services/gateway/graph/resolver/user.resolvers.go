package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yuorei/hackathon/go/user"
	"github.com/yuorei/hackathon/graph/generated"
	"github.com/yuorei/hackathon/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserPayload, error) {
	response, err := r.client.CreateUser(input)
	if err != nil {
		return nil, err
	}

	// id が 0 の場合はエラー
	if response.Id == 0 {
		// emailの中にエラーが入っている
		return nil, fmt.Errorf("failed to create user: %s", response.Email)
	}

	userID := strconv.Itoa(int(response.Id))
	userID = fmt.Sprintf("user_%s", userID)

	var gender model.Gender
	switch response.Gender {
	case user.Gender_MAN:
		gender = model.GenderMan
	case user.Gender_WOMAN:
		gender = model.GenderWoman
	case user.Gender_GENDER_OTHER:
		gender = model.GenderOther
	}

	var affiliation model.Affiliation
	switch response.Affiliation {
	case user.Affiliation_STUDENT:
		affiliation = model.AffiliationStudent
	case user.Affiliation_AFFILIATION_OTHER:
		affiliation = model.AffiliationOther
	}

	return &model.UserPayload{
		ID:          userID,
		Name:        response.Name,
		Email:       response.Email,
		Gender:      &gender,
		Affiliation: &affiliation,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.UserPayload, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
