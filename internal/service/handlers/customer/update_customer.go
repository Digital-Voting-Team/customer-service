package handlers

import (
	"customer-service/internal/data"
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/customer"
	"customer-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateCustomerRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	customer, err := helpers.CustomersQ(r).FilterByID(request.CustomerID).Get()
	if customer == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newCustomer := data.Customer{
		RegistrationDate: &request.Data.Attributes.RegistrationDate,
		PersonID:         cast.ToInt64(request.Data.Relationships.Person.Data.ID),
	}

	relatePerson, err := helpers.PersonsQ(r).FilterByID(newCustomer.PersonID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new person")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var resultCustomer data.Customer
	resultCustomer, err = helpers.CustomersQ(r).FilterByID(customer.ID).Update(newCustomer)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update customer")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Person{
		Key: resources.NewKeyInt64(relatePerson.ID, resources.PERSON),
		Attributes: resources.PersonAttributes{
			Name:     relatePerson.Name,
			Phone:    relatePerson.Phone,
			Email:    relatePerson.Email,
			Birthday: relatePerson.Birthday,
		},
		Relationships: resources.PersonRelationships{
			Address: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(relatePerson.AddressID, 10),
					Type: resources.ADDRESS,
				},
			},
		},
	})

	result := resources.CustomerResponse{
		Data: resources.Customer{
			Key: resources.NewKeyInt64(resultCustomer.ID, resources.CUSTOMER),
			Attributes: resources.CustomerAttributes{
				RegistrationDate: *resultCustomer.RegistrationDate,
			},
			Relationships: resources.CustomerRelationships{
				Person: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultCustomer.PersonID, 10),
						Type: resources.PERSON,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
