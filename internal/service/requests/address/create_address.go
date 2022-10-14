package requests

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	"github.com/Digital-Voting-Team/customer-service/resources"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateAddressRequest struct {
	Data resources.Address
}

func NewCreateAddressRequest(r *http.Request) (CreateAddressRequest, error) {
	var request CreateAddressRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *CreateAddressRequest) validate() error {
	return mergeErrors(validation.Errors{
		"/data/attributes/building_number": validation.Validate(&r.Data.Attributes.BuildingNumber, validation.Required,
			validation.By(helpers.IsInteger)),
		"/data/attributes/street": validation.Validate(&r.Data.Attributes.Street, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/city": validation.Validate(&r.Data.Attributes.City, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/district": validation.Validate(&r.Data.Attributes.District, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/region": validation.Validate(&r.Data.Attributes.Region, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/postal_code": validation.Validate(&r.Data.Attributes.PostalCode, validation.Required,
			validation.Length(1, 45)),
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
