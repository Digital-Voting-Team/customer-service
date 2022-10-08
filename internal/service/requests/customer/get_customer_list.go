package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetPersonListRequest struct {
	pgdb.OffsetPageParams
	FilterName   []string `filter:"name"`
	FilterPhone  []string `filter:"phone"`
	FilterEmails []string `filter:"emails"`
}

func NewGetPersonListRequest(r *http.Request) (GetPersonListRequest, error) {
	var request GetPersonListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
