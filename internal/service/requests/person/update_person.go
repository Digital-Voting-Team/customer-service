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

type UpdatePersonRequest struct {
	PersonID int64 `url:"-" json:"-"`
	Data     resources.Person
}

func NewUpdatePersonRequest(r *http.Request) (UpdatePersonRequest, error) {
	request := UpdatePersonRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.PersonID = cast.ToInt64(chi.URLParam(r, "id"))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *UpdatePersonRequest) validate() error {
	return mergeErrors(validation.Errors{
		"/data/attributes/name": validation.Validate(&r.Data.Attributes.Name, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/phone": validation.Validate(&r.Data.Attributes.Phone, validation.Required,
			validation.Length(3, 30)),
		"/data/attributes/email": validation.Validate(&r.Data.Attributes.Email, validation.Required,
			validation.Length(3, 45)),
		"/data/attributes/birthday": validation.Validate(&r.Data.Attributes.Birthday,
			validation.By(helpers.IsDate)),
		"/data/relationships/address/data/id": validation.Validate(&r.Data.Relationships.Address.Data.ID,
			validation.Required, validation.By(helpers.IsInteger)),
	}).Filter()
}
