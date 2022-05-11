package sanitizer

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
	"reflect"
)

type Sanitizer interface {
	SanitizeString(fieldName, fieldValue string) string
}

type TypeValue struct {
	FieldType   reflect.StructField
	StructField reflect.StructField
	FieldValue  reflect.Value
}

func SanitizeInput(source any, sanitizer Sanitizer) any {

	sourceRef := TypeValue{
		FieldType:  reflect.TypeOf(source).Field(0),
		FieldValue: reflect.ValueOf(source),
	}

	sourceValue := reflect.ValueOf(source)
	sourceCopy := reflect.New(sourceValue.Type()).Elem()

	sanitize(sourceRef, sourceCopy, sanitizer)

	return sourceCopy.Interface()
}

func sanitize(source TypeValue, target reflect.Value, sanitizer Sanitizer) {

	switch source.FieldValue.Kind() {

	case reflect.String:
		log.Info("call sanitizer")
		log.Info("current field",
			zap.Any("name", source.FieldType.Name),
			zap.Any("type", source.FieldType.Type),
			zap.Any("value", source.FieldValue.String()),
		)

		target.SetString(sanitizer.SanitizeString(source.FieldType.Name, source.FieldValue.String()))

	case reflect.Struct:
		for i := 0; i < source.FieldValue.NumField(); i += 1 {
			embeddedTypeValue := TypeValue{
				FieldType:  source.FieldValue.Type().Field(i),
				FieldValue: source.FieldValue.Field(i),
			}
			sanitize(embeddedTypeValue, target.Field(i), sanitizer)
		}

	case reflect.Pointer:
		sourceValue := source.FieldValue.Elem()

		// Check if the pointer is nil
		if !sourceValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		target.Set(reflect.New(sourceValue.Type()))

		sourceType := source.FieldType
		sourceType.Type = sourceValue.Type()

		extractedTypeValue := TypeValue{
			FieldType:  sourceType,
			FieldValue: sourceValue,
		}

		sanitize(extractedTypeValue, target.Elem(), sanitizer)

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

		sanitize(extractedTypeValue, targetValue, sanitizer)
		target.Set(targetValue)
	}
}
