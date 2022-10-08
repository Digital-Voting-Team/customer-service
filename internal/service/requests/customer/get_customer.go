package requests

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cast"

	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetCustomerRequest struct {
	CustomerID int64 `url:"-"`
}

func NewGetCustomerRequest(r *http.Request) (GetCustomerRequest, error) {
	request := GetCustomerRequest{}

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.CustomerID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
