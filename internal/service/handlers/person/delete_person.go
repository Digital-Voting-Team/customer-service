package handlers

import (
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/customer-service/internal/service/requests/person"
	staffRes "github.com/Digital-Voting-Team/staff-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
	if accessLevel < staffRes.Manager {
		helpers.Log(r).Info("insufficient user permissions")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

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
