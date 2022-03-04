package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"net/http"
)

const (
	AddContactsEndpoint = "/v3/marketing/contacts"
)

type SendgridClient struct {
	config Config
}

func NewSendgridClient(config *Config) *SendgridClient {
	return &SendgridClient{config: *config}
}

func (s *SendgridClient) AddContacts(ctx context.Context, contacts []Contact) (*ImportContactsResponse, error) {

	req := sendgrid.GetRequest(s.config.ApiKey, AddContactsEndpoint, s.config.Host)
	req.Method = http.MethodPut

	requestData, err := json.Marshal(AddContactsRequest{
		ListIds:  s.config.ContactListIds,
		Contacts: contacts,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AddContactsRequest: %v", err)
	}

	req.Body = requestData

	var result *ImportContactsResponse
	err = s.makeRequest(ctx, AddContactsEndpoint, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SendgridClient) makeRequest(ctx context.Context, endpoint string, request rest.Request, response interface{}) error {

	res, err := sendgrid.MakeRequestWithContext(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to make request to the sendgrid api, endpoint: %q, error: %v", endpoint, err)
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("sendgrid api returned an error, endpoint: %q, status: %q, body: %q", endpoint, res.StatusCode, res.Body)
	}

	err = json.Unmarshal([]byte(res.Body), response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal sendgrid response, endpoint: %q, error: %v", endpoint, err)
	}

	return nil
}
