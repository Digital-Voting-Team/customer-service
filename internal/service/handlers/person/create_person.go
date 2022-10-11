package handlers

import (
	"customer-service/internal/data"
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/person"
	"customer-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreatePersonRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	person := data.Person{
		Name:      request.Data.Attributes.Name,
		Phone:     request.Data.Attributes.Phone,
		Email:     request.Data.Attributes.Email,
		Birthday:  request.Data.Attributes.Birthday,
		AddressID: cast.ToInt64(request.Data.Relationships.Address.Data.ID),
	}

	var resultPerson data.Person
	relateAddress, err := helpers.AddressesQ(r).FilterByID(person.AddressID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get address")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultPerson, err = helpers.PersonsQ(r).Insert(person)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create person")
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
