package sanitizer

import (
	"fmt"
	"github.com/volatiletech/null/v8"
	"reflect"
)

type FieldSanitizer interface {
	SanitizeString(field reflect.StructField, value string) (string, error)
	SanitizeNullString(field reflect.StructField, value null.String) (null.String, error)
}

type TypeValue struct {
	FieldType  *reflect.StructField
	FieldValue *reflect.Value
}

const TagName = "sanitize"

// SanitizeInput traverses a given struct and applies the currently set Sanitizer to any field of type string or
// null.String.
//
// Example:
//
//	type MyInput struct {
//		MyTitle string `sanitize:"strict"`
//		OptionalDescription null.String `sanitize:"html"`
//	}
//
//	myInput := MyInput{
//		MyTitle:             "<h1>Title</h1>",
//		OptionalDescription: null.StringFrom("<p><a href=\"javascript:alert('XSS')\">Some Description</a></p>"),
//	}
//
//	_ = sanitizer.SanitizeInput(&myInput)
//
//	fmt.Println(res.MyTitle) // Title
//	fmt.Println(res.OptionalDescription.String) // <p>Some Description</p>
//
// See Sanitizer on the default setting and how to override those. For applying a custom struct for just the given
// input use SanitizeInputWithLocalSanitizer.
//
// SanitizeInput expects a non-nil pointer to a struct, otherwise it will return an error. Note that due to map indexes
// not being addressable passing any kind of embedded map types will fails as well.
func SanitizeInput(obj any) error {
	return SanitizeInputWithLocalSanitizer(obj, Sanitizer)
}

// MustSanitizeInput behaves like SanitizeInput but panics instead of returning an error.
func MustSanitizeInput(obj any) {
	err := SanitizeInputWithLocalSanitizer(obj, Sanitizer)
	if err != nil {
		panic(err)
	}
}

// SanitizeInputWithLocalSanitizer can be used to sanitize some input with a custom sanitizer. This can be useful
// if you have input that you want to treat differently than other input.
//
// To apply a custom sanitizer globally you can just override Sanitizer instead.
func SanitizeInputWithLocalSanitizer(obj any, sanitizer FieldSanitizer) error {

	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("invalid parameter given expected non-nil pointer to a struct got %q instead", reflect.TypeOf(obj)))
	}

	sourceValue := reflect.ValueOf(obj).Elem()
	sourceType := sourceValue.Type().Field(0)

	sourceRef := TypeValue{
		FieldType:  &sourceType,
		FieldValue: &sourceValue,
	}

	return sanitize(sourceRef, sanitizer)
}

// MustSanitizeInputWithLocalSanitizer behaves like SanitizeInputWithLocalSanitizer but panics instead of returning an error.
func MustSanitizeInputWithLocalSanitizer(obj any, sanitizer FieldSanitizer) {
	err := SanitizeInputWithLocalSanitizer(obj, sanitizer)
	if err != nil {
		panic(err)
	}
}

// sanitize traverses data structures recursively and replaces the value of all string or null.String fields with the
// results of the sanitizer.
//
// To pass on the field tags we use the TypeValue here containing both the fields reflect.Value and reflect.StructField.
// When we dereference values (such as pointers or interfaces) we use a little hack here, we just pass the pointer's
// reflect.StructField and overwrite the Type there with the one of the dereferenced field.
func sanitize(source TypeValue, sanitizer FieldSanitizer) error {

	switch source.FieldValue.Kind() {

	case reflect.String:
		// if it's a string we just replace it with the result of the sanitizer.
		sanitizedField, err := sanitizer.SanitizeString(*source.FieldType, source.FieldValue.String())
		if err != nil {
			return err
		}
		source.FieldValue.SetString(sanitizedField)

	case reflect.Struct:
		// if this struct is of type null.String we just replace it with the sanitizer's result as well
		if source.FieldValue.CanConvert(reflect.TypeOf(null.String{})) {
			sanitizedField, err := sanitizer.SanitizeNullString(*source.FieldType, source.FieldValue.Interface().(null.String))
			if err != nil {
				return err
			}
			source.FieldValue.Set(reflect.ValueOf(sanitizedField))
		} else {
			// otherwise, we traverse through all the struct fields and apply the sanitizer on them as well
			for i := 0; i < source.FieldValue.NumField(); i += 1 {

				elementValue := source.FieldValue.Field(i)
				elementType := source.FieldValue.Type().Field(i)

				embeddedTypeValue := TypeValue{
					FieldType:  &elementType,
					FieldValue: &elementValue,
				}
				err := sanitize(embeddedTypeValue, sanitizer)
				if err != nil {
					return err
				}
			}
		}

	case reflect.Pointer:
		sourceValue := source.FieldValue.Elem()

		// Break here on nil pointers
		if !sourceValue.IsValid() {
			return nil
		}

		sourceValue.Type()
		sourceType := source.FieldType
		sourceType.Type = sourceValue.Type()

		extractedTypeValue := TypeValue{
			FieldType:  sourceType,
			FieldValue: &sourceValue,
		}

		err := sanitize(extractedTypeValue, sanitizer)
		if err != nil {
			return err
		}

	case reflect.Interface:
		sourceValue := source.FieldValue.Elem()
		sourceType := source.FieldType
		sourceType.Type = sourceValue.Type()

		extractedTypeValue := TypeValue{
			FieldType:  sourceType,
			FieldValue: &sourceValue,
		}

		err := sanitize(extractedTypeValue, sanitizer)
		if err != nil {
			return err
		}

	case reflect.Slice:
		for i := 0; i < source.FieldValue.Len(); i += 1 {

			elementValue := source.FieldValue.Index(i)
			elementType := source.FieldType
			elementType.Type = elementValue.Type()

			elementTypeValue := TypeValue{
				FieldType:  elementType,
				FieldValue: &elementValue,
			}

			err := sanitize(elementTypeValue, sanitizer)
			if err != nil {
				return err
			}
		}

	case reflect.Map:
		// Unfortunately in the current design we are not able to sanitize map values. This is as map index operations
		// are not addressable, which means we are not able to replace values within a map (we only receive a copy of
		// the map values but require a pointer).
		return fmt.Errorf("cannot sanitize map inputs as they are not addressable")
	}

	return nil
}
