package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type DeleteCustomerRequest struct {
	CustomerID int64 `url:"-"`
}

func NewDeleteCustomerRequest(r *http.Request) (DeleteCustomerRequest, error) {
	request := DeleteCustomerRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CustomerID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
