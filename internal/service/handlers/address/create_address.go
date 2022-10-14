package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/data"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/address"
	"github.com/Digital-Voting-Team/customer-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateAddressRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var resultAddress data.Address

	address := data.Address{
		BuildingNumber: request.Data.Attributes.BuildingNumber,
		Street:         request.Data.Attributes.Street,
		City:           request.Data.Attributes.City,
		District:       request.Data.Attributes.District,
		Region:         request.Data.Attributes.Region,
		PostalCode:     request.Data.Attributes.PostalCode,
	}

	resultAddress, err = helpers.AddressesQ(r).Insert(address)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create address")
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
