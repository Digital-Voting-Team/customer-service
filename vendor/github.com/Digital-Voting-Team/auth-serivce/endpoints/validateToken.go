package endpoints

import (
	"encoding/json"
	"github.com/Digital-Voting-Team/auth-serivce/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func ParseJwtResponse(r *http.Response) (*resources.JwtResponse, error) {
	var response resources.JwtResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal JwtResponse")
	}

	return &response, nil
}

func ValidateToken(token, endpoint string) (*resources.JwtResponse, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build new request")
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to send request, endpoint: "+endpoint)
	}

	return ParseJwtResponse(res)
}
