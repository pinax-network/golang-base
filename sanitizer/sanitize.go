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
	FieldType  reflect.StructField
	FieldValue reflect.Value
}

const TagName = "sanitize"

// SanitizeInput traverses a struct and applies the currently set Sanitizer to any field of type string or
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
//	res, _ := sanitizer.SanitizeInput(myInput)
//
//	fmt.Println(res.MyTitle) // Title
//	fmt.Println(res.OptionalDescription.String) // <p>Some Description</p>
//
// See Sanitizer on the default setting and how to override those. For applying a custom struct for just the given
// input use SanitizeInputWithLocalSanitizer.
func SanitizeInput[T any](source T) (res T, err error) {
	return SanitizeInputWithLocalSanitizer(source, Sanitizer)
}

// MustSanitizeInput behaves like SanitizeInput but panics instead of returning an error.
func MustSanitizeInput[T any](source T) (res T) {
	res, err := SanitizeInputWithLocalSanitizer(source, Sanitizer)
	if err != nil {
		panic(err)
	}

	return res
}

// SanitizeInputWithLocalSanitizer can be used to sanitize some input with a custom sanitizer. This can be useful
// if you have input that you want to treat differently than other input.
//
// To apply a custom sanitizer globally you can just override Sanitizer instead.
func SanitizeInputWithLocalSanitizer[T any](source T, sanitizer FieldSanitizer) (res T, err error) {

	if reflect.TypeOf(source).Kind() == reflect.Struct {

		sourceRef := TypeValue{
			FieldType:  reflect.TypeOf(source).Field(0),
			FieldValue: reflect.ValueOf(source),
		}

		sourceValue := reflect.ValueOf(source)
		sourceCopy := reflect.New(sourceValue.Type()).Elem()

		err = sanitize(sourceRef, sourceCopy, sanitizer)
		if err != nil {
			return
		}

		res = sourceCopy.Interface().(T)
		return
	}

	err = fmt.Errorf("invalid type %q given as source, the source needs to be a struct", reflect.TypeOf(source))
	return
}

func MustSanitizeInputWithLocalSanitizer[T any](source T, sanitizer FieldSanitizer) T {

	res, err := SanitizeInputWithLocalSanitizer(source, sanitizer)
	if err != nil {
		panic(err)
	}

	return res
}

func sanitize(source TypeValue, target reflect.Value, sanitizer FieldSanitizer) error {

	switch source.FieldValue.Kind() {

	case reflect.String:
		sanitizedField, err := sanitizer.SanitizeString(source.FieldType, source.FieldValue.String())
		if err != nil {
			return err
		}
		target.SetString(sanitizedField)

	case reflect.Struct:

		// if this struct is a kind of null.String we sanitize it as well, otherwise we traverse further
		if source.FieldValue.CanConvert(reflect.TypeOf(null.String{})) {
			sanitizedField, err := sanitizer.SanitizeNullString(source.FieldType, source.FieldValue.Interface().(null.String))
			if err != nil {
				return err
			}
			target.Set(reflect.ValueOf(sanitizedField))
		} else {
			for i := 0; i < source.FieldValue.NumField(); i += 1 {
				embeddedTypeValue := TypeValue{
					FieldType:  source.FieldValue.Type().Field(i),
					FieldValue: source.FieldValue.Field(i),
				}
				err := sanitize(embeddedTypeValue, target.Field(i), sanitizer)
				if err != nil {
					return err
				}
			}
		}

	case reflect.Pointer:
		sourceValue := source.FieldValue.Elem()

		// Check if the pointer is nil
		if !sourceValue.IsValid() {
			return nil
		}
		// Allocate a new object and set the pointer to it
		target.Set(reflect.New(sourceValue.Type()))

		sourceType := source.FieldType
		sourceType.Type = sourceValue.Type()

		extractedTypeValue := TypeValue{
			FieldType:  sourceType,
			FieldValue: sourceValue,
		}

		err := sanitize(extractedTypeValue, target.Elem(), sanitizer)
		if err != nil {
			return err
		}

	case reflect.Interface:
		sourceValue := source.FieldValue.Elem()

		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		targetValue := reflect.New(sourceValue.Type()).Elem()

		sourceType := source.FieldType
		sourceType.Type = sourceValue.Type()

		extractedTypeValue := TypeValue{
			FieldType:  sourceType,
			FieldValue: sourceValue,
		}

		err := sanitize(extractedTypeValue, targetValue, sanitizer)
		if err != nil {
			return err
		}

		target.Set(targetValue)

	case reflect.Slice:
		target.Set(reflect.MakeSlice(source.FieldType.Type, source.FieldValue.Len(), source.FieldValue.Cap()))
		for i := 0; i < source.FieldValue.Len(); i += 1 {

			elementValue := source.FieldValue.Index(i)
			elementType := source.FieldType
			elementType.Type = elementValue.Type()

			elementTypeValue := TypeValue{
				FieldType:  elementType,
				FieldValue: elementValue,
			}

			err := sanitize(elementTypeValue, target.Index(i), sanitizer)
			if err != nil {
				return err
			}
		}

	case reflect.Map:
		target.Set(reflect.MakeMap(source.FieldType.Type))
		for _, key := range source.FieldValue.MapKeys() {
			elementValue := source.FieldValue.MapIndex(key)
			// New gives us a pointer, but again we want the value
			targetValue := reflect.New(elementValue.Type()).Elem()

			elementType := source.FieldType
			elementType.Type = elementValue.Type()

			elementTypeValue := TypeValue{
				FieldType:  elementType,
				FieldValue: elementValue,
			}

			err := sanitize(elementTypeValue, targetValue, sanitizer)
			if err != nil {
				return err
			}

			target.SetMapIndex(key, targetValue)
		}

	default:
		target.Set(source.FieldValue)
	}

	return nil
}
