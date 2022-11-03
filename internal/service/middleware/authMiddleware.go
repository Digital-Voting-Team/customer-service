package middleware

import (
	"context"
	"github.com/Digital-Voting-Team/auth-serivce/endpoints"
	"github.com/Digital-Voting-Team/customer-service/internal/config"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func BasicAuth(endpointsConf *config.EndpointsConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtResponse, err := endpoints.ValidateToken(
				r.Header.Get("Authorization"),
				endpointsConf.Endpoints["auth_service"],
			)
			if err != nil || jwtResponse.Data.ID == "" {
				helpers.Log(r).WithError(err).Info("auth failed")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			ctx := context.WithValue(r.Context(), "jwt", jwtResponse)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
