package shufti

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/helper"
	"github.com/eosnationftw/eosn-base-api/response"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ShuftiAuthMiddleware struct {
	secret string
}

func NewShuftiAuthMiddleware(config *Config) (*ShuftiAuthMiddleware, error) {

	if config.Secret == "" {
		return nil, errors.New("no shufti secret set in the config")
	}

	return &ShuftiAuthMiddleware{secret: config.Secret}, nil
}

func (s *ShuftiAuthMiddleware) VerifySignature() gin.HandlerFunc {
	return func(c *gin.Context) {

		signature := c.GetHeader("Signature")

		if signature == "" {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "no signature given")
			return
		}

		rawBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			helper.ReportPrivateErrorAndAbort(c, response.InternalServerError, fmt.Errorf("failed to read request body: %v", err))
			return
		}

		// write back request body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))

		if !s.verifySignatureHeader(signature, string(rawBody)) {
			helper.ReportPublicErrorAndAbort(c, response.Unauthorized, "invalid signature given")
			return
		}

		c.Next()
	}
}

func (s *ShuftiAuthMiddleware) verifySignatureHeader(signature, requestBody string) bool {

	// calculate sha256sum(request_body + shufti_secret)
	checksum := sha256.Sum256([]byte(requestBody + s.secret))

	// valid if hex representation of the checksum equals the given signature
	isValid := fmt.Sprintf("%x", checksum) == signature

	return isValid
}
