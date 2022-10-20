package requests

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	"github.com/Digital-Voting-Team/customer-service/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateCustomerRequest struct {
	Data resources.Customer
}

func NewCreateCustomerRequest(r *http.Request) (CreateCustomerRequest, error) {
	var request CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateCustomerRequest) validate() error {
	return helpers.MergeErrors(validation.Errors{
		"/data/attributes/registration_date": validation.Validate(&r.Data.Attributes.RegistrationDate,
			validation.Required, validation.By(helpers.IsDate)),
		"/data/relationships/person/data/id": validation.Validate(&r.Data.Relationships.Person.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
		"/data/relationships/person/user/id": validation.Validate(&r.Data.Relationships.User.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
