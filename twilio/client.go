package twilio

import (
	"github.com/friendsofgo/errors"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type Client struct {
	config *Config
	client *twilio.RestClient
}

var (
	ErrFailedToRequestCode = errors.New("failed to request code")
	ErrFailedToVerifyCode  = errors.New("failed to verify code")
	ErrCodeNotApproved     = errors.New("code not approved")
)

func NewClient(config *Config) (*Client, error) {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   config.ApiKey,
		Password:   config.ApiSecret,
		AccountSid: config.AccountSID,
	})

	return &Client{config: config, client: client}, nil
}

func (c *Client) RequestCode(phoneNumber string, channel string) error {

	params := &openapi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel(channel)

	_, err := c.client.VerifyV2.CreateVerification(c.config.VerifyServiceSID, params)
	if err != nil {
		incRequestErrorCounter()
		return errors.WithMessage(ErrFailedToRequestCode, err.Error())
	}
	incRequestCounter()

	return nil
}

func (c *Client) VerifyCode(phoneNumber string, code string) error {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := c.client.VerifyV2.CreateVerificationCheck(c.config.VerifyServiceSID, params)
	if err != nil {
		incVerifyErrorCounter()
		return errors.WithMessage(ErrFailedToVerifyCode, err.Error())
	} else if *resp.Status != "approved" {
		return errors.WithMessage(ErrCodeNotApproved, "verification status: "+*resp.Status)
	}
	incVerifyCounter()

	return nil
}
