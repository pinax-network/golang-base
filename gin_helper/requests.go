package gin_helper

import (
	"fmt"
	"github.com/eosnationftw/eosn-base-api/sanitizer"
	"github.com/gin-gonic/gin"
	"reflect"
)

// BindAndSanitizeJSON binds the JSON request body of a gin.Context to the given struct and applies the HTML sanitizer
// on it to strip malicious HTML. This requires the given obj struct to have 'sanitize' tags applied. See
// sanitizer.SanitizeInput for more information about this tag.
//
// The obj parameter is expected as a non-nil pointer to a struct, otherwise it will panic. The error return value
// contains errors from binding and validating the JSON request body.
func BindAndSanitizeJSON(c *gin.Context, obj any) error {

	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		panic(fmt.Errorf("invalid parameter given: expected non-nil pointer got: %q instead", reflect.TypeOf(obj)))
	}

	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}

	sanitizer.MustSanitizeInput(obj)

	return nil
}
