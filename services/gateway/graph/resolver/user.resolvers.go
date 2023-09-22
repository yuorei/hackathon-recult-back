package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yuorei/hackathon/go/image"
	"github.com/yuorei/hackathon/go/user"
	"github.com/yuorei/hackathon/graph/generated"
	"github.com/yuorei/hackathon/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserPayload, error) {
	var responseImage *image.UploadResponse
	var imageURL *string
	if input.ProfileImage != nil {
		var err error
		responseImage, err = r.client.UploadImage(input.ProfileImage)
		if err != nil {
			return nil, err
		}
		imageURL = &responseImage.ImageUrl
	}
	responseUser, err := r.client.CreateUser(input, imageURL)
	if err != nil {
		return nil, err
	}

	// id が 0 の場合はエラー
	if responseUser.Id == 0 {
		// emailの中にエラーが入っている
		return nil, fmt.Errorf("failed to create user: %s", responseUser.Email)
	}

	userID := strconv.Itoa(int(responseUser.Id))
	userID = fmt.Sprintf("user_%s", userID)

	var gender model.Gender
	switch responseUser.Gender {
	case user.Gender_MAN:
		gender = model.GenderMan
	case user.Gender_WOMAN:
		gender = model.GenderWoman
	default:
		gender = model.GenderOther
	}

	var affiliation model.Affiliation
	switch responseUser.Affiliation {
	case user.Affiliation_STUDENT:
		affiliation = model.AffiliationStudent
	default:
		affiliation = model.AffiliationOther
	}

	return &model.UserPayload{
		ID:              userID,
		Name:            responseUser.Name,
		Email:           responseUser.Email,
		ProfileImageURL: imageURL,
		Gender:          &gender,
		Affiliation:     &affiliation,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.UserPayload, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
