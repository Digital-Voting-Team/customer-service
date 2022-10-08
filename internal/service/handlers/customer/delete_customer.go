package handlers

import (
	"customer-service/internal/service/helpers"
	requests "customer-service/internal/service/requests/person"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeletePersonRequest(r)
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

	err = helpers.PersonsQ(r).Delete(request.PersonID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete person")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
