package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
	"html"

	sentry "github.com/getsentry/sentry-go"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model/graph"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/graph/generated"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input graph.NewPost) (*model.Post, error) {
	escapedText := html.EscapeString(input.Text)
	created, err := r.PostUseCase.CreatePost(&escapedText, input.UserID)
	if err != nil {
		err = fmt.Errorf("resolver CreatePost() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return created, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.NewUser) (*model.User, error) {
	escapedName := html.EscapeString(input.Name)
	escapedEmail := html.EscapeString(input.Email)
	escapedPassword := html.EscapeString(input.Password)
	created, err := r.UserUseCase.CreateUser(&escapedName, &escapedEmail, &escapedPassword)
	if err != nil {
		err = fmt.Errorf("resolver CreateUser() err %w", err)
		sentry.CaptureException(err)
		return nil, err
	}

	return created, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
