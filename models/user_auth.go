package models

import (
	"context"
	"net/http"
	"pizza-backend/jwt"

	"gorm.io/gorm"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type ClaimUser struct {
	Id    int
	Name  string
	Email string
	Roles []string
}

func (c *ClaimUser) CheckRoles(roles []string) bool {
	return true
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" {
				ctx := context.WithValue(r.Context(), userCtxKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			bearer := "Bearer "
			auth = auth[len(bearer):]

			validate, err := jwt.JwtValidate(context.Background(), auth)
			if err != nil || !validate.Valid {
				ctx := context.WithValue(r.Context(), userCtxKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			customClaim, _ := validate.Claims.(*jwt.JwtCustomClaim)

			session := &Session{}
			err = db.Where("user_id = ?", customClaim.ID).Find(&session).Error
			if err != nil {
				ctx := context.WithValue(r.Context(), userCtxKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			var user *User
			err = db.Where("id = ?", customClaim.ID).First(&user).Error
			if err != nil {
				ctx := context.WithValue(r.Context(), userCtxKey, nil)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			var u = ClaimUser{Id: user.ID, Name: user.Name, Email: user.Email, Roles: user.Roles.Strings()}
			ctx := context.WithValue(r.Context(), userCtxKey, &u)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ClaimUser {
	raw, _ := ctx.Value(userCtxKey).(*ClaimUser)
	return raw
}
