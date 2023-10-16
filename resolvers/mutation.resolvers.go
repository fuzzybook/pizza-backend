package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"pizza-backend/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	return models.CreateUser(ctx, input)
}

// UpdateUserRoles is the resolver for the updateUserRoles field.
func (r *mutationResolver) UpdateUserRoles(ctx context.Context, input models.UpdateUserRoles) (*models.User, error) {
	return models.UpdateRoles(ctx, input)
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input models.UpdateUserPassword) (*models.User, error) {
	return models.UpdatePassword(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input models.UserLogin) (*models.Session, error) {
	return models.Login(ctx, input)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
