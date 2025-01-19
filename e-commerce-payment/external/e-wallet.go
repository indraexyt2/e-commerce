package external

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"e-commerce-payment/helpers"
	"e-commerce-payment/internal/models"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type PaymentLinkResponse struct {
	Message string `json:"message"`
	Data    struct {
		OTP string `json:"otp"`
	} `json:"data"`
}

func (e *External) generateSignature(ctx context.Context, payload, timestamp, endpoint string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	payload = re.ReplaceAllString(payload, "")
	payload = strings.ToLower(payload) + timestamp + endpoint

	secretKey := helpers.GetEnv("WALLET_SECRET_KEY")
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func (e *External) PaymentLink(ctx context.Context, req *models.PaymentMethodLink) (*PaymentLinkResponse, error) {
	url := helpers.GetEnv("WALLET_URL") + helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_LINK")

	reqMap := map[string]int{
		"wallet_id": req.SourceID,
	}
	bytePayload, err := json.Marshal(reqMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bytePayload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	timestamp := time.Now().Format(time.RFC3339)
	signature := e.generateSignature(ctx, string(bytePayload), timestamp, helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_LINK"))

	httpReq.Header.Set("Client-ID", helpers.GetEnv("WALLET_CLIENT_ID"))
	httpReq.Header.Set("Timestamp", timestamp)
	httpReq.Header.Set("Signature", signature)

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to link payment")
	}

	response := &PaymentLinkResponse{}
	err = json.NewDecoder(httpRes.Body).Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}

func (e *External) PaymentUnlink(ctx context.Context, req *models.PaymentMethodLink) (*PaymentLinkResponse, error) {
	url := helpers.GetEnv("WALLET_URL") + fmt.Sprintf(helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_UNLINK"), req.SourceID)

	httpReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	timestamp := time.Now().Format(time.RFC3339)
	signature := e.generateSignature(ctx, "", timestamp, fmt.Sprintf(helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_UNLINK"), req.SourceID))

	httpReq.Header.Set("Client-ID", helpers.GetEnv("WALLET_CLIENT_ID"))
	httpReq.Header.Set("Timestamp", timestamp)
	httpReq.Header.Set("Signature", signature)

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to unlink payment")
	}

	response := &PaymentLinkResponse{}
	err = json.NewDecoder(httpRes.Body).Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}

func (e *External) PaymentLinkConfirmation(ctx context.Context, sourceID int, otp string) (*PaymentLinkResponse, error) {
	url := helpers.GetEnv("WALLET_URL") + fmt.Sprintf(helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_LINK_CONFIRM"), sourceID)

	reqMap := map[string]string{
		"otp": otp,
	}
	bytePayload, err := json.Marshal(reqMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}
	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(bytePayload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	timestamp := time.Now().Format(time.RFC3339)
	signature := e.generateSignature(ctx, string(bytePayload), timestamp, fmt.Sprintf(helpers.GetEnv("WALLET_ENDPOINT_PAYMENT_LINK_CONFIRM"), sourceID))

	httpReq.Header.Set("Client-ID", helpers.GetEnv("WALLET_CLIENT_ID"))
	httpReq.Header.Set("Timestamp", timestamp)
	httpReq.Header.Set("Signature", signature)

	client := &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return nil, errors.New("failed to link confirmation")
	}

	response := &PaymentLinkResponse{}
	err = json.NewDecoder(httpRes.Body).Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}
