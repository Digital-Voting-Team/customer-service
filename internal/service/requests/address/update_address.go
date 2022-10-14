package requests

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/customer-service/internal/service/helpers"
	"github.com/Digital-Voting-Team/customer-service/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateAddressRequest struct {
	AddressID int64 `url:"-" json:"-"`
	Data      resources.Address
}

func NewUpdateAddressRequest(r *http.Request) (UpdateAddressRequest, error) {
	request := UpdateAddressRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.AddressID = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateAddressRequest) validate() error {
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
