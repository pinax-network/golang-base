package sanitizer

import (
	"reflect"
)

type Sanitizer interface {
	SanitizeString(fieldName, fieldValue string) string
}

type TypeValue struct {
	FieldType  reflect.StructField
	FieldValue reflect.Value
}

// SanitizeInput traverses any arbitary struct, applies the given sanitizer on all fields of type string and returns a
// deep copy of the given struct with it's fields sanitized.
// Note: this method panics if the given source is not a struct.
func SanitizeInput[T any](source T, sanitizer Sanitizer) T {

	if reflect.TypeOf(source).Kind() == reflect.Struct {

		sourceRef := TypeValue{
			FieldType:  reflect.TypeOf(source).Field(0),
			FieldValue: reflect.ValueOf(source),
		}

		sourceValue := reflect.ValueOf(source)
		sourceCopy := reflect.New(sourceValue.Type()).Elem()

		sanitize(sourceRef, sourceCopy, sanitizer)

		return sourceCopy.Interface().(T)
	}

	panic("invalid type given, needs to be a struct")
}

func sanitize(source TypeValue, target reflect.Value, sanitizer Sanitizer) {

	switch source.FieldValue.Kind() {

	case reflect.String:
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

			sanitize(elementTypeValue, target.Index(i), sanitizer)
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

			sanitize(elementTypeValue, targetValue, sanitizer)
			target.SetMapIndex(key, targetValue)
		}

	default:
		target.Set(source.FieldValue)
	}
}
