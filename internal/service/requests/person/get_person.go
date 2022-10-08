package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetPersonRequest struct {
	PersonID int64 `url:"-"`
}

func NewGetPersonRequest(r *http.Request) (GetPersonRequest, error) {
	request := GetPersonRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.PersonID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
