package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	base_response "github.com/eosnationftw/eosn-base-api/response"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

// ReadResponseBody reads the body of the given response without modifying it and tries to unmarshal it into the given target interface.
func ReadResponseBody(response *http.Response, target interface{}) error {

	// read the current body
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	// put the body content back as the response body is now empty
	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// unmarshal into the general api response format
	var apiResponse base_response.ApiDataResponse
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal into general api response: %v", err)
	}

	// decode the api response data into the given target interface
	err = mapstructure.Decode(apiResponse.Data, target)
	if err != nil {
		return fmt.Errorf("failed to decode api response data field into given interface: %v", err)
	}

	return nil
}

// BasicAuth creates the base64 encoded string containing user:password for basic authentication headers
func BasicAuth(user, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))
}
