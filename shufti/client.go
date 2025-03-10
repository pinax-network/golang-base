package shufti

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pinax-network/golang-base/helper"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	RequestPending            = "request.pending"
	RequestInvalid            = "request.invalid"
	VerificationCancelled     = "verification.cancelled"
	RequestTimeout            = "request.timeout"
	RequestUnauthorized       = "request.unauthorized"
	VerificationAccepted      = "verification.accepted"
	VerificationDeclined      = "verification.declined"
	VerificationStatusChanged = "verification.status.changed"
	RequestDeleted            = "request.deleted"
	RequestReceived           = "request.received"
)

type Client struct {
	httpClient *http.Client
	config     *Config
}

func NewClient(config *Config) (*Client, error) {

	if !strings.HasPrefix(config.Host, "https") {
		return nil, fmt.Errorf("TLS required for shufti host url, %q given", config.Host)
	}

	res := &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		config:     config,
	}

	return res, nil
}

func (c *Client) GetVerificationUrlTtl() int {
	return c.config.VerificationUrlTtl
}

func (c *Client) GetVerificationUrl(reference, email, redirectUrl string) (string, error) {

	requestData := getDefaultDocumentVerificationRequest()
	requestData.Reference = reference
	requestData.Email = email
	requestData.Ttl = c.config.VerificationUrlTtl
	requestData.RedirectUrl = redirectUrl

	if c.config.CallbackUrl != "" {
		requestData.CallbackUrl = c.config.CallbackUrl
	}

	requestBytes, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.config.Host, bytes.NewBuffer(requestBytes))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", helper.BasicAuth(c.config.ClientId, c.config.Secret)))
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 300 {
		return "", fmt.Errorf("error on shuftipro api: %s", string(bodyData))
	}

	var result VerificationResponse
	err = json.Unmarshal(bodyData, &result)
	if err != nil {
		return "", err
	}

	return result.VerificationUrl, nil
}
