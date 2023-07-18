package auth

import (
	"context"
	"log"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"shopolahUser"}

type contextKey struct {
	name string
}

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	View
)

// A stand-in for our database backed user object
type UserViewer struct {
	Role Role
}

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool // If viewer is admin.
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

/*
FIXME: `context` is not being updated.
*/
func MiddlewareTest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			newCtx := context.WithValue(r.Context(), userCtxKey, UserViewer{Role: View})

			r = r.WithContext(newCtx)
			viewer := FromContext(r.Context())
			if viewer != nil {
				log.Println("viewer: ", *viewer)
			} else {
				log.Println("viewer: ", viewer)
			}

			next.ServeHTTP(w, r)

		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func FromContext(ctx context.Context) *UserViewer {
	raw, _ := ctx.Value(userCtxKey).(*UserViewer)
	return raw
}
