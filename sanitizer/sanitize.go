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
			zap.Any("value", source.FieldValue.String()),
		)

		target.SetString(sanitizer.SanitizeString(source.FieldType.Name, source.FieldValue.String()))

	case reflect.Struct:
		for i := 0; i < source.FieldValue.NumField(); i += 1 {

			varName := source.FieldValue.Type().Field(i).Name
			varType := source.FieldValue.Type().Field(i).Type
			varValue := source.FieldValue.Field(i).Interface()

			log.Info("struct field",
				zap.Any("varName", varName),
				zap.Any("varType", varType),
				zap.Any("varValue", varValue),
			)

			embeddedTypeValue := TypeValue{
				FieldType:  source.FieldValue.Type().Field(i),
				FieldValue: source.FieldValue.Field(i),
			}
			sanitize(embeddedTypeValue, target.Field(i), sanitizer)
		}
	}
}
