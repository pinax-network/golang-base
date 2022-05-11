package sanitizer

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
	"testing"
)

type LogSanitizer struct {
	t *testing.T
}

func (l LogSanitizer) SanitizeString(fieldName, fieldValue string) string {
	log.Info("call SanitizeString()", zap.String("fieldName", fieldName), zap.String("fieldValue", fieldValue))

	return fieldValue
}

type TestStruct struct {
	SomeField string
}

func TestSum(t *testing.T) {

	_ = log.InitializeLogger(true)

	logSanitizer := LogSanitizer{}
	testSource := TestStruct{SomeField: "test"}

	res := SanitizeInput(testSource, logSanitizer)

	log.Info("result", zap.Any("res", res))
}
