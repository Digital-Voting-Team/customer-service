package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeletePersonRequest struct {
	PersonID int64 `url:"-"`
}

func NewDeletePersonRequest(r *http.Request) (DeletePersonRequest, error) {
	request := DeletePersonRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.PersonID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
