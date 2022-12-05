package middleware

import (
	"context"
	authEndoints "github.com/Digital-Voting-Team/auth-service/endpoints"
	"github.com/Digital-Voting-Team/customer-service/internal/config"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	staffEndoints "github.com/Digital-Voting-Team/staff-service/endpoints"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func BasicAuth(endpointsConf *config.EndpointsConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtResponse, err := authEndoints.ValidateToken(
				r.Header.Get("Authorization"),
				endpointsConf.Endpoints["auth-service"],
			)
			if jwtResponse == nil {
				helpers.Log(r).WithError(err).Info("auth failed, jwtResponse == nil", endpointsConf.Endpoints["auth-service"])
				ape.Render(w, problems.BadRequest(err))
				return
			}
			if err != nil || jwtResponse.Data.ID == "" {
				helpers.Log(r).WithError(err).Info("auth failed")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			positionResponse, err := staffEndoints.GetPosition(
				r.Header.Get("Authorization"),
				endpointsConf.Endpoints["staff-service"],
			)
			if positionResponse == nil {
				helpers.Log(r).Info("auth failed, positionResponse == nil")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			if err != nil || positionResponse.Data.ID == "" {
				helpers.Log(r).WithError(err).Info("unable to get position (auth)")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			ctx := context.WithValue(r.Context(), "accessLevel", positionResponse.Data.Attributes.AccessLevel)
			ctx = context.WithValue(ctx, "userId", cast.ToInt64(jwtResponse.Data.Relationships.User.Data.ID))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
