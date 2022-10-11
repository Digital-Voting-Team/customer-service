package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetAddressListRequest struct {
	pgdb.OffsetPageParams
	FilterBuildingNumber []int64  `filter:"building_number"`
	FilterStreet         []string `filter:"street"`
	FilterCity           []string `filter:"city"`
	FilterDistrict       []string `filter:"district"`
	FilterRegion         []string `filter:"region"`
	FilterPostalCode     []string `filter:"postal_code"`
}

func NewGetAddressListRequest(r *http.Request) (GetAddressListRequest, error) {
	var request GetAddressListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
