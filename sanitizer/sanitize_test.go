package sanitizer

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
	"testing"
)

type LogSanitizer struct {
	t *testing.T
}

type TestSanitizer struct{}

func (l LogSanitizer) SanitizeString(fieldName, fieldValue string) string {
	log.Info("call SanitizeString()", zap.String("fieldName", fieldName), zap.String("fieldValue", fieldValue))

	return fieldValue
}

func (t TestSanitizer) SanitizeString(fieldName, fieldValue string) string {
	return fieldValue + "_sanitized"
}

type TestStruct struct {
	StringField    string
	StringFieldPtr *string
}

func TestSum(t *testing.T) {

	_ = log.InitializeLogger(true)

	logSanitizer := TestSanitizer{}

	testStringPtr := "test_string_pointer"

	testSource := TestStruct{
		StringField:    "test_string_field",
		StringFieldPtr: &testStringPtr,
	}

	res := SanitizeInput(testSource, logSanitizer)

	log.Info("result", zap.Any("res", res))
}
