package dfuse

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type DfuseClient struct {
	GraphQLClient GraphQLClient
}

var plainTextDialOption = grpc.WithInsecure()
var insecureTLSDialOption = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
var secureTLSDialOption = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
var maxCallRecvMsgSize = 1024 * 1024 * 100
var defaultCallOptions = []grpc.CallOption{
	grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize),
	grpc.WaitForReady(true),
}

var keepaliveDialOption = grpc.WithKeepaliveParams(keepalive.ClientParameters{
	// Send pings every (x seconds) there is no activity
	Time: 30 * time.Second,
	// Wait that amount of time for ping ack before considering the connection dead
	Timeout: 10 * time.Second,
	// Send pings even without active streams
	PermitWithoutStream: true,
})

func NewDfuseClient(endpoint string, insecure bool) (*DfuseClient, error) {

	res := &DfuseClient{}

	dialOpts := []grpc.DialOption{keepaliveDialOption, grpc.WithDefaultCallOptions(defaultCallOptions...)}
	if insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		transportCreds := credentials.NewClientTLSFromCert(nil, "")
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(transportCreds))
	}

	dfuseAPIKey := os.Getenv("DFUSE_API_KEY")
	if dfuseAPIKey != "" {
		token, _, err := getToken(dfuseAPIKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get dfuse token: %w", err)
		}

		credential := oauth.NewOauthAccess(&oauth2.Token{AccessToken: token, TokenType: "Bearer"})
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(credential))
	}

	conn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %w", err)
	}

	res.GraphQLClient = NewGraphQLClient(conn)

	return res, nil
}

func getToken(apiKey string) (token string, expiration time.Time, err error) {
	reqBody := bytes.NewBuffer([]byte(fmt.Sprintf(`{"api_key":"%s"}`, apiKey)))
	resp, err := http.Post("https://auth.eosnation.io/v1/auth/issue", "application/json", reqBody)
	if err != nil {
		err = fmt.Errorf("unable to obtain token: %s", err)
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("unable to obtain token, status not 200, got %d: %s", resp.StatusCode, reqBody.String())
		return
	}

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		token = gjson.GetBytes(body, "token").String()
		expiration = time.Unix(gjson.GetBytes(body, "expires_at").Int(), 0)
	}
	return
}
