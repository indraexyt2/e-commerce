package external

import (
	"bytes"
	"context"
	"e-commerce-payment/helpers"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type OrderResponse struct {
	Message string `json:"message"`
}

func (e *External) OrderCallback(ctx context.Context, orderID int, status string) (*OrderResponse, error) {
	url := helpers.GetEnv("E_COMMERCE_URL") + fmt.Sprintf(helpers.GetEnv("E_COMMERCE_ENDPOINT_ORDER_CALLBACK"), orderID)
	fmt.Println("URL: ", url)
	reqMap := map[string]string{
		"status": status,
	}
	bytePayload, err := json.Marshal(reqMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}
	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(bytePayload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to add transaction")
	}

	response := &OrderResponse{}
	err = json.NewDecoder(httpRes.Body).Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}
