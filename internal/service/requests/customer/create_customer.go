package requests

import (
	"customer-service/resources"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
	return mergeErrors(validation.Errors{
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.CreatedAt, validation.Required,
			validation.Length(3, 45)),
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
