package handlers

import (
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/person"
	"customer-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPersonRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultPerson, err := helpers.PersonsQ(r).FilterByID(request.PersonID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get person from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultPerson == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relateAddress, err := helpers.AddressesQ(r).FilterByID(resultPerson.AddressID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get address")
		ape.RenderErr(w, problems.NotFound())
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
