package service

import (
	"customer-service/internal/data/pg"
	"customer-service/internal/service/helpers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	log := s.log.WithFields(map[string]interface{}{
		"service": "customer-service-api",
	})

	r.Use(
		ape.RecoverMiddleware(log),
		ape.LoganMiddleware(log),
		ape.CtxMiddleware(
			helpers.CtxLog(log),
			helpers.CtxAddressesQ(pg.NewAddressesQ(s.db)),
			helpers.CtxPersonsQ(pg.NewPersonsQ(s.db)),
			helpers.CtxCustomersQ(pg.NewCustomersQ(s.db)),
		),
	)
	r.Route("/integrations/customer-service", func(r chi.Router) {
		// configure endpoints here
	})

	return r
}
