package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/data"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/person"
	"github.com/Digital-Voting-Team/customer-service/resources"
	staffRes "github.com/Digital-Voting-Team/staff-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdatePersonRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	person, err := helpers.PersonsQ(r).FilterByID(request.PersonID).Get()
	if person == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	userId := r.Context().Value("userId").(int64)
	accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
	_, personId, _, err := helpers.GetIdsForGivenUser(r, userId)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong relations")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if accessLevel != staffRes.Admin && personId != person.ID {
		helpers.Log(r).Info("insufficient user permissions")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	newPerson := data.Person{
		Name:      request.Data.Attributes.Name,
		Phone:     request.Data.Attributes.Phone,
		Email:     request.Data.Attributes.Email,
		Birthday:  request.Data.Attributes.Birthday,
		AddressID: cast.ToInt64(request.Data.Relationships.Address.Data.ID),
	}

	relateAddress, err := helpers.AddressesQ(r).FilterByID(newPerson.AddressID).Get()
	if err != nil || relateAddress == nil {
		helpers.Log(r).WithError(err).Error("failed to get new address")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultPersonByAddress, err := helpers.PersonsQ(r).FilterByAddressID(person.AddressID).Get()
	if resultPersonByAddress.ID == 0 || resultPersonByAddress.AddressID != newPerson.AddressID {
		helpers.Log(r).WithError(err).Error("invalid address to update")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	var resultPerson data.Person
	resultPerson, err = helpers.PersonsQ(r).FilterByID(person.ID).Update(newPerson)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update person")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Address{
		Key: resources.NewKeyInt64(relateAddress.ID, resources.ADDRESS),
		Attributes: resources.AddressAttributes{
			BuildingNumber: relateAddress.BuildingNumber,
			Street:         relateAddress.Street,
			City:           relateAddress.City,
			District:       relateAddress.District,
			Region:         relateAddress.Region,
			PostalCode:     relateAddress.PostalCode,
		},
	})

	result := resources.PersonResponse{
		Data: resources.Person{
			Key: resources.NewKeyInt64(resultPerson.ID, resources.PERSON),
			Attributes: resources.PersonAttributes{
				Name:     resultPerson.Name,
				Phone:    resultPerson.Phone,
				Email:    resultPerson.Email,
				Birthday: resultPerson.Birthday,
			},
			Relationships: resources.PersonRelationships{
				Address: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultPerson.AddressID, 10),
						Type: resources.ADDRESS,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
