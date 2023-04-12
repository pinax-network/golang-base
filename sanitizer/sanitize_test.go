package sanitizer

import (
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"reflect"
	"testing"
)

type TestFieldSanitizer struct{}

func (t TestFieldSanitizer) SanitizeString(field reflect.StructField, value string) (string, error) {
	if fieldTag := field.Tag.Get(TagName); fieldTag == "test" {
		return t.SanitizeStringField(value), nil
	}

	return "", fmt.Errorf("invalid sanitizer")
}

func (t TestFieldSanitizer) SanitizeNullString(field reflect.StructField, value null.String) (null.String, error) {
	if !value.Valid {
		return value, nil
	}

	sanitizedString, err := t.SanitizeString(field, value.String)
	if err != nil {
		return value, err
	}

	return null.StringFrom(sanitizedString), nil
}

func (t TestFieldSanitizer) SanitizeStringField(value string) string {
	return value + "_sanitized"
}

func init() {
	Sanitizer = TestFieldSanitizer{}
}

func TestStringField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField string `sanitize:"test"`
	}{
		TestField: testField,
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	testStruct.TestField = testSanitizer.SanitizeStringField(testField)
	assert.Equal(t, testStruct, res)
}

func TestStringPtrField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField *string `sanitize:"test"`
	}{
		TestField: &testField,
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	sanitizedField := testSanitizer.SanitizeStringField(testField)
	testStruct.TestField = &sanitizedField
	assert.Equal(t, testStruct, res)
}

func TestStringSliceField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testStruct := struct {
		TestField []string `sanitize:"test"`
	}{
		TestField: []string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	for i, f := range testStruct.TestField {
		testStruct.TestField[i] = testSanitizer.SanitizeStringField(f)
	}

	assert.Equal(t, testStruct, res)
}

func TestStringPtrSliceField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestField []*string `sanitize:"test"`
	}{
		TestField: []*string{&testString1, &testString2, &testString3},
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	for i, f := range testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeStringField(*f)
		testStruct.TestField[i] = &sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

func TestStringSlicePtrField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testStruct := struct {
		TestField *[]string `sanitize:"test"`
	}{
		TestField: &[]string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	for i, f := range *testStruct.TestField {
		(*testStruct.TestField)[i] = testSanitizer.SanitizeStringField(f)
	}

	assert.Equal(t, testStruct, res)
}

func TestStringPtrSlicePtrField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestField *[]*string `sanitize:"test"`
	}{
		TestField: &[]*string{&testString1, &testString2, &testString3},
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	for i, f := range *testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeStringField(*f)
		(*testStruct.TestField)[i] = &sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

func TestMapField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestMap map[string]string `sanitize:"test"`
	}{
		TestMap: map[string]string{
			"field1": testString1,
			"field2": testString2,
			"field3": testString3,
		},
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)

	for i, f := range testStruct.TestMap {
		sanitizedString := testSanitizer.SanitizeStringField(f)
		testStruct.TestMap[i] = sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

type StringType string

func TestStringTypeField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	var testField StringType
	testField = "test_field"
	testStruct := struct {
		TestField StringType `sanitize:"test"`
	}{
		TestField: testField,
	}

	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)
	testStruct.TestField = StringType(testSanitizer.SanitizeStringField(string(testField)))
	assert.Equal(t, testStruct, res)
}

func TestEmbeddedStruct(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	type embeddedTestStruct struct {
		TestField string `sanitize:"test"`
	}

	testField := "test_field"
	testStruct := struct {
		Embedded embeddedTestStruct `sanitize:"dive"`
	}{
		embeddedTestStruct{TestField: testField},
	}
	res, err := SanitizeInput(testStruct)
	require.NoError(t, err)
	testStruct.Embedded.TestField = testSanitizer.SanitizeStringField(testField)
	assert.Equal(t, testStruct, res)
}

func TestLocalSanitizer(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField string `sanitize:"test"`
	}{
		TestField: testField,
	}

	res, err := SanitizeInputWithLocalSanitizer(testStruct, testSanitizer)
	require.NoError(t, err)

	testStruct.TestField = testSanitizer.SanitizeStringField(testField)
	assert.Equal(t, testStruct, res)
}

func TestInvalidSanitizer(t *testing.T) {

	testField := "test_field"
	testStruct := struct {
		TestField string `sanitize:"asdf"`
	}{
		TestField: testField,
	}

	_, err := SanitizeInput(testStruct)
	assert.Equal(t, err.Error(), "invalid sanitizer")
}
