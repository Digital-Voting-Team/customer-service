package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/address"
	staffRes "github.com/Digital-Voting-Team/staff-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
	if accessLevel < staffRes.Manager {
		helpers.Log(r).Info("insufficient user permissions")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	request, err := requests.NewDeleteAddressRequest(r)
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

	err = helpers.AddressesQ(r).Delete(request.AddressID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
