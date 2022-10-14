package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/customer"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteCustomerRequest(r)
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

	err = helpers.CustomersQ(r).Delete(request.CustomerID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete customer")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
