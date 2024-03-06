package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pinax-network/golang-base/helper"
	"github.com/pinax-network/golang-base/log"
	"github.com/pinax-network/golang-base/response"
	"github.com/pinax-network/golang-base/validate"
	"go.uber.org/zap"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Run all pre-handler middleware & handlers before handling errors.
		c.Next()

		// todo go through all errors or only use last?
		err := c.Errors.Last()
		if err == nil {
			return
		}

		httpErr, ok := err.Err.(*response.ApiError)
		if !ok {
			log.Error("error middleware received error that is not of type *ApiError", zap.Any("error", err))
			httpErr = response.InternalServerError
		}

		// log internal server errors in detail
		if httpErr.Is(response.InternalServerError) {
			user, _ := helper.ExtractUserFromContext(c)
			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			log.Error("[Internal server error]",
				zap.Time("time", time.Now()),
				zap.Any("error", err),
				zap.Any("meta", err.Meta),
				zap.Any("user", user),
				zap.String("request", string(httpRequest)),
			)
		}

		resultErrs := []*response.ApiError{httpErr}

		// if we have a non-private error or we are in debug mode try to get the error detail from parsing the meta interface
		if (!err.IsType(gin.ErrorTypePrivate) || gin.IsDebugging()) && err.Meta != nil {
			metaInterface := err.Meta
			switch v := metaInterface.(type) {
			case string:
				httpErr.Detail = v
				break
			case validator.ValidationErrors:
				resultErrs = parseValidationErrors(httpErr, v)
				break
			case *validate.SortValidationError:
				resultErrs = parseSortValidationErrors(httpErr, v)
				break
			case *strconv.NumError:
				httpErr.Detail = fmt.Sprintf("failed to parse value '%s': %s", v.Num, v.Err.Error())
				break
			case error:
				httpErr.Detail = v.Error()
				break
			default:
				log.Error("received unknown type for error meta", zap.Any("meta_interface", metaInterface))
			}
		}

		errResponse := response.ApiErrorResponse{Errors: resultErrs}
		c.AbortWithStatusJSON(httpErr.Status, errResponse)
	}
}

func parseValidationErrors(apiError *response.ApiError, fieldErrors []validator.FieldError) (apiErrors []*response.ApiError) {

	if len(fieldErrors) == 0 {
		apiErrors = []*response.ApiError{apiError}
		return
	}

	apiErrors = make([]*response.ApiError, len(fieldErrors))

	for i, fieldErr := range fieldErrors {
		var sb strings.Builder
		sb.WriteString("validation failed on field '" + fieldErr.Field() + "'")
		sb.WriteString(", condition: " + fieldErr.ActualTag())
		// Print condition parameters, e.g. oneof=red blue -> { red blue }
		if fieldErr.Param() != "" {
			sb.WriteString(" { " + fieldErr.Param() + " }")
		}
		if fieldErr.Value() != nil && fieldErr.Value() != "" {
			sb.WriteString(fmt.Sprintf(", actual: %v", fieldErr.Value()))
		}

		apiErrors[i] = &response.ApiError{
			Status: apiError.Status,
			Code:   apiError.Code,
			Detail: sb.String(),
		}
	}

	return
}

func parseSortValidationErrors(apiError *response.ApiError, ve *validate.SortValidationError) (apiErrors []*response.ApiError) {

	apiErrors = make([]*response.ApiError, len(ve.Errors))

	for i, e := range ve.Errors {
		apiErrors[i] = &response.ApiError{
			Status: apiError.Status,
			Code:   apiError.Code,
			Detail: e.Error(),
		}
	}
	return
}
