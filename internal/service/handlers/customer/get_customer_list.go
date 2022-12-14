package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/data"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/customer"
	"github.com/Digital-Voting-Team/customer-service/resources"
	staffRes "github.com/Digital-Voting-Team/staff-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetCustomerList(w http.ResponseWriter, r *http.Request) {
	accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
	if accessLevel < staffRes.Manager {
		helpers.Log(r).Info("insufficient user permissions")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	request, err := requests.NewGetCustomerListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	customersQ := helpers.CustomersQ(r)
	applyFilters(customersQ, request)
	customers, err := customersQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get customers")
		ape.Render(w, problems.InternalError())
		return
	}
	persons, err := helpers.PersonsQ(r).FilterByID(getPersonsIDs(customers)...).Select()

	response := resources.CustomerListResponse{
		Data:     newCustomerList(customers),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newCustomerIncluded(persons),
	}
	ape.Render(w, response)
}

func applyFilters(q data.CustomersQ, request requests.GetCustomerListRequest) {
	q.Page(request.OffsetPageParams)

	if request.FilterDateAfter != nil {
		q.FilterByDateAfter(*request.FilterDateAfter)
	}

	if request.FilterDateBefore != nil {
		q.FilterByDateBefore(*request.FilterDateBefore)
	}
}

func newCustomerList(customers []data.Customer) []resources.Customer {
	result := make([]resources.Customer, len(customers))
	for i, customer := range customers {
		result[i] = resources.Customer{
			Key: resources.NewKeyInt64(customer.ID, resources.CUSTOMER),
			Attributes: resources.CustomerAttributes{
				RegistrationDate: *customer.RegistrationDate,
			},
			Relationships: resources.CustomerRelationships{
				Person: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(customer.PersonID, 10),
						Type: resources.PERSON,
					},
				},
				User: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(customer.UserID, 10),
						Type: resources.USER_REF,
					},
				},
			},
		}
	}
	return result
}

func getPersonsIDs(customers []data.Customer) []int64 {
	personIDs := make([]int64, len(customers))
	for i := 0; i < len(customers); i++ {
		personIDs[i] = customers[i].PersonID
	}
	return personIDs
}

func newCustomerIncluded(persons []data.Person) resources.Included {
	result := resources.Included{}
	for _, item := range persons {
		resource := newPersonModel(item)
		result.Add(&resource)
	}
	return result
}

func newPersonModel(person data.Person) resources.Person {
	return resources.Person{
		Key: resources.NewKeyInt64(person.ID, resources.PERSON),
		Attributes: resources.PersonAttributes{
			Name:     person.Name,
			Phone:    person.Phone,
			Email:    person.Email,
			Birthday: person.Birthday,
		},
		Relationships: resources.PersonRelationships{
			Address: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(person.AddressID, 10),
					Type: resources.ADDRESS,
				},
			},
		},
	}
}
