package external

import (
	"context"
	"e-commerce-order/helpers"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type Profile struct {
	Message string `json:"message"`
	Data    struct {
		Username    string `json:"username"`
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		Dob         string `json:"dob"`
		Role        string `json:"role"`
	} `json:"data"`
}

type External struct{}

func (ext *External) GetProfile(ctx context.Context, token string) (*Profile, error) {
	url := helpers.GetEnv("UMS_URL") + helpers.GetEnv("UMS_ENDPOINT_PROFILE")

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	httpReq.Header.Set("Authorization", token)

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get profile")
	}

	response := &Profile{}
	err = json.NewDecoder(httpRes.Body).Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}
