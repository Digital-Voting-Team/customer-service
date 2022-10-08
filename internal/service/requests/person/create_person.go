package requests

import (
	"customer-service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreatePersonRequest struct {
	Data resources.Person
}

func NewCreatePersonRequest(r *http.Request) (CreatePersonRequest, error) {
	var request CreatePersonRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreatePersonRequest) validate() error {
	return mergeErrors(validation.Errors{
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.Name, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/phone": validation.Validate(&r.Data.Attributes.Phone, validation.Required,
			validation.Length(3, 30)),
		"/data/attributes/email": validation.Validate(&r.Data.Attributes.Email, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/address_id": validation.Validate(&r.Data.Attributes.Email, validation.Required),
	}).Filter()
}

func mergeErrors(validationErrors ...validation.Errors) validation.Errors {
	result := make(validation.Errors)
	for _, errs := range validationErrors {
		for key, err := range errs {
			result[key] = err
		}
	}
	return result
}
