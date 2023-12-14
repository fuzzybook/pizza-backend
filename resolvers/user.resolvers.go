package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"pizza-backend/models"
)

// User is the resolver for the User field.
func (r *sessionResolver) User(ctx context.Context, obj *models.Session) (*models.User, error) {
	return models.GetCompleteUserById(ctx, obj.ID)
}

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *models.User) ([]models.UserRole, error) {
	return obj.Roles, nil
}

// Session returns SessionResolver implementation.
func (r *Resolver) Session() SessionResolver { return &sessionResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type sessionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
