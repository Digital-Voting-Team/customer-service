package handlers

import (
	"customer-service/internal/data"
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/address"
	"customer-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateAddressRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address, err := helpers.AddressesQ(r).FilterByID(request.AddressID).Get()
	if address == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	var resultAddress data.Address

	newAddress := data.Address{
		BuildingNumber: request.Data.Attributes.BuildingNumber,
		Street:         request.Data.Attributes.Street,
		City:           request.Data.Attributes.City,
		District:       request.Data.Attributes.District,
		Region:         request.Data.Attributes.Region,
		PostalCode:     request.Data.Attributes.PostalCode,
	}

	resultAddress, err = helpers.AddressesQ(r).Update(newAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.AddressResponse{
		Data: resources.Address{
			Key: resources.NewKeyInt64(resultAddress.ID, resources.ADDRESS),
			Attributes: resources.AddressAttributes{
				BuildingNumber: resultAddress.BuildingNumber,
				Street:         resultAddress.Street,
				City:           resultAddress.City,
				District:       resultAddress.District,
				Region:         resultAddress.Region,
				PostalCode:     resultAddress.PostalCode,
			},
		},
	}
	ape.Render(w, result)
}