package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/data"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/customer"
	"github.com/Digital-Voting-Team/customer-service/resources"
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

	//jwt := r.Context().Value("jwt").(resources_auth.JwtResponse)
	//if request.Data.Relationships.User.Data.ID != jwt.Data.Relationships.User.Data.ID {
	//	helpers.Log(r).WithError(err).Info("jwt user is inconsistent with request user")
	//	ape.RenderErr(w, problems.BadRequest(err)...)
	//	return
	//}

	newCustomer := data.Customer{
		RegistrationDate: &request.Data.Attributes.RegistrationDate,
		PersonID:         cast.ToInt64(request.Data.Relationships.Person.Data.ID),
		UserID:           cast.ToInt64(request.Data.Relationships.User.Data.ID),
	}

	relatePerson, err := helpers.PersonsQ(r).FilterByID(newCustomer.PersonID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new person")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultCustomerByUser, err := helpers.CustomersQ(r).FilterByUserID(customer.UserID).Get()
	if (resultCustomerByUser != nil || resultCustomerByUser.ID != 0) &&
		resultCustomerByUser.UserID != newCustomer.UserID {
		helpers.Log(r).WithError(err).Error("User already related to customer")
		ape.RenderErr(w, problems.Conflict())
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
				User: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultCustomer.UserID, 10),
						Type: resources.USER_REF,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
