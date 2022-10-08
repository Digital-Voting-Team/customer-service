package requests

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetCustomerListRequest struct {
	pgdb.OffsetPageParams
	FilterDateAfter  *time.Time `filter:"date_after"`
	FilterDateBefore *time.Time `filter:"date_before"`
}

func NewGetCustomerListRequest(r *http.Request) (GetCustomerListRequest, error) {
	var request GetCustomerListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
