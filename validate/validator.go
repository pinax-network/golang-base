package validate

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

type JsonValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *JsonValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr && value.IsValid() {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *JsonValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *JsonValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		err := v.validate.RegisterValidation("notblank", validators.NotBlank)
		log.FatalIfError("failed to initialize 'notblank' validation", err)

		err = v.validate.RegisterValidation("sortpair", func(fl validator.FieldLevel) bool {
			regex := regexp.MustCompile("^([A-Za-z_]+)(:(asc|desc))?$")
			return regex.MatchString(fl.Field().String())
		})
		log.FatalIfError("failed to initialize 'sortpair' validation", err)

		err = v.validate.RegisterValidation("eosaccount", func(fl validator.FieldLevel) bool {
			regex := regexp.MustCompile("^[a-z1-5.]{1,11}[a-z1-5]$")
			return regex.MatchString(fl.Field().String())
		})
		log.FatalIfError("failed to initialize 'eosaccount' validation", err)

		err = v.validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
			regex := regexp.MustCompile("^[a-z\\d]([a-z\\d]|\\.([a-z\\d])){2,38}$")
			return regex.MatchString(fl.Field().String())
		})
		log.FatalIfError("failed to initialize 'username' validation", err)
	})
}
