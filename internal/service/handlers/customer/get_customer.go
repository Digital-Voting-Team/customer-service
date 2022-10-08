package handlers

import (
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/customer"
	"customer-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetCustomerRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultCustomer, err := helpers.CustomersQ(r).FilterByID(request.CustomerID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get customer from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultCustomer == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relatePerson, err := helpers.PersonsQ(r).FilterByID(resultCustomer.PersonID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get person")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Person{
		Key: resources.NewKeyInt64(relatePerson.ID, resources.PERSON),
		Attributes: resources.PersonAttributes{
			Name:  relatePerson.Name,
			Phone: relatePerson.Phone,
			Email: relatePerson.Email,
		},
		Relationships: resources.PersonRelationships{
			Address: resources.Relation{
				Data: &resources.Key{
					ID: strconv.FormatInt(relatePerson.AddressID, 10),
				},
			},
		},
	})

	result := resources.CustomerResponse{
		Data: resources.Customer{
			Key: resources.NewKeyInt64(resultCustomer.ID, resources.CUSTOMER),
			Attributes: resources.CustomerAttributes{
				CreatedAt: *resultCustomer.CreatedAt,
			},
			Relationships: resources.CustomerRelationships{
				Person: resources.Relation{
					Data: &resources.Key{
						ID: strconv.FormatInt(resultCustomer.PersonID, 10),
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
