package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/address"
	"github.com/Digital-Voting-Team/customer-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetAddressRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address, err := helpers.AddressesQ(r).FilterByID(request.AddressID).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get address from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if address == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	result := resources.AddressResponse{
		Data: resources.Address{
			Key: resources.NewKeyInt64(address.ID, resources.ADDRESS),
			Attributes: resources.AddressAttributes{
				BuildingNumber: address.BuildingNumber,
				Street:         address.Street,
				City:           address.City,
				District:       address.District,
				Region:         address.Region,
				PostalCode:     address.PostalCode,
			},
		},
	}

	ape.Render(w, result)
}
