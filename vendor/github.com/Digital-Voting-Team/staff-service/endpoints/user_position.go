package endpoints

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/staff-service/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type UserPositionRequest struct {
	UserId int64
}

func ParsePositionResponse(r *http.Response) (*resources.PositionResponse, error) {
	var response resources.PositionResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal")
	}

	return &response, nil
}

func GetPosition(token, endpoint string) (*resources.PositionResponse, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return ParsePositionResponse(res)
}
