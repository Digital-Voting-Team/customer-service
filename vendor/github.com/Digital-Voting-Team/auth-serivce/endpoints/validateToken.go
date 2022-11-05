package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/Digital-Voting-Team/auth-serivce/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func ParseJwtResponse(r *http.Response) (*resources.JwtResponse, error) {
	var response resources.JwtResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response, errors.Wrap(err, "failed to unmarshal")
	}

	return &response, nil
}

func ValidateToken(token, endpoint string) (*resources.JwtResponse, error) {
	fmt.Println(endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return ParseJwtResponse(res)
}
