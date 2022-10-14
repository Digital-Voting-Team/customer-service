package service

import (
	"github.com/Digital-Voting-Team/customer-service/internal/data/pg"
	address "github.com/Digital-Voting-Team/customer-service/internal/service/handlers/address"
	customer "github.com/Digital-Voting-Team/customer-service/internal/service/handlers/customer"
	person "github.com/Digital-Voting-Team/customer-service/internal/service/handlers/person"

	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"

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
		r.Route("/addresses", func(r chi.Router) {
			r.Post("/", address.CreateAddress)
			r.Get("/", address.GetAddressList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", address.GetAddress)
				r.Put("/", address.UpdateAddress)
				r.Delete("/", address.DeleteAddress)
			})
		})
		r.Route("/persons", func(r chi.Router) {
			r.Post("/", person.CreatePerson)
			r.Get("/", person.GetPersonList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", person.GetPerson)
				r.Put("/", person.UpdatePerson)
				r.Delete("/", person.DeletePerson)
			})
		})
		r.Route("/customers", func(r chi.Router) {
			r.Post("/", customer.CreateCustomer)
			r.Get("/", customer.GetCustomerList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", customer.GetCustomer)
				r.Put("/", customer.UpdateCustomer)
				r.Delete("/", customer.DeleteCustomer)
			})
		})
	})

	return r
}
