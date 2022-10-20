package middleware

import (
	"context"
	"github.com/Digital-Voting-Team/auth-serivce/endpoints"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"os"
)

func BasicAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtResponse, err := endpoints.ValidateToken(
				r.Header.Get("Authorization"),
				os.Getenv("AUTH_SERVICE"),
			)
			if err != nil || jwtResponse.Data.ID == "" {
				ape.Render(w, problems.BadRequest(err))
				return
			}
			ctx := context.WithValue(r.Context(), "jwt", jwtResponse)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
