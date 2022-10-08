package requests

import (
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
		"/data/attributes/created_at": validation.Validate(&r.Data.Attributes.CreatedAt, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/person_id": validation.Validate(&r.Data.Relationships.Person.Data.ID, validation.Required),
	}).Filter()
}
