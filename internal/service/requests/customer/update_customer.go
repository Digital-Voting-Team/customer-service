package requests

import (
	"customer-service/internal/service/helpers"
	"customer-service/resources"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type UpdateCustomerRequest struct {
	CustomerID int64 `url:"-" json:"-"`
	Data       resources.Customer
}

func NewUpdateCustomerRequest(r *http.Request) (UpdateCustomerRequest, error) {
	request := UpdateCustomerRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CustomerID = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdateCustomerRequest) validate() error {
	return mergeErrors(validation.Errors{
		"/data/attributes/registration_date": validation.Validate(&r.Data.Attributes.RegistrationDate,
			validation.Required, validation.By(helpers.IsDate)),
		"/data/relationships/person/data/id": validation.Validate(&r.Data.Relationships.Person.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
