package handlers

import (
	"customer-service/internal/data"
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/person"
	"customer-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPersonList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPersonListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	personsQ := helpers.PersonsQ(r)
	applyFilters(personsQ, request)
	persons, err := personsQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get persons")
		ape.Render(w, problems.InternalError())
		return
	}
	addresses, err := helpers.AddressesQ(r).FilterByID(getAddressIDs(persons)...).Select()

	response := resources.PersonListResponse{
		Data:     newPersonsList(persons),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newPersonIncluded(addresses),
	}
	ape.Render(w, response)
}

func applyFilters(q data.PersonsQ, request requests.GetPersonListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterName) > 0 {
		q.FilterByNames(request.FilterName...)
	}

	if len(request.FilterPhone) > 0 {
		q.FilterByPhones(request.FilterPhone...)
	}

	if len(request.FilterEmails) > 0 {
		q.FilterByEmails(request.FilterEmails...)
	}
}

func newPersonsList(persons []data.Person) []resources.Person {
	result := make([]resources.Person, len(persons))
	for i, person := range persons {
		result[i] = resources.Person{
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
	return result
}

func getAddressIDs(persons []data.Person) []int64 {
	addressIDs := make([]int64, len(persons))
	for i := 0; i < len(persons); i++ {
		addressIDs[i] = persons[i].AddressID
	}
	return addressIDs
}

func newPersonIncluded(addresses []data.Address) resources.Included {
	result := resources.Included{}
	for _, item := range addresses {
		resource := newAddressModel(item)
		result.Add(&resource)
	}
	return result
}

func newAddressModel(address data.Address) resources.Address {
	return resources.Address{
		Key: resources.NewKeyInt64(address.ID, resources.ADDRESS),
		Attributes: resources.AddressAttributes{
			BuildingNumber: address.BuildingNumber,
			Street:         address.Street,
			City:           address.City,
			District:       address.District,
			Region:         address.Region,
			PostalCode:     address.PostalCode,
		},
	}
}
