package sanitizer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	res := testSanitizer.SanitizeStringField(testField)

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
}

func TestStringPtrField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField *string `sanitize:"test"`
	}{
		TestField: &testField,
	}

	res := testSanitizer.SanitizeStringField(testField)

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, &res)
}

func TestStringSliceField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testStruct := struct {
		TestField []string `sanitize:"test"`
	}{
		TestField: []string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res := make([]string, len(testStruct.TestField))
	for i, f := range testStruct.TestField {
		res[i] = testSanitizer.SanitizeStringField(f)
	}

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
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

	res := make([]*string, len(testStruct.TestField))
	for i, f := range testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeStringField(*f)
		res[i] = &sanitizedString
	}

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
}

func TestStringSlicePtrField(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testStruct := struct {
		TestField *[]string `sanitize:"test"`
	}{
		TestField: &[]string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res := make([]string, len(*testStruct.TestField))
	for i, f := range *testStruct.TestField {
		res[i] = testSanitizer.SanitizeStringField(f)
	}

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, &res)
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

	res := make([]*string, len(*testStruct.TestField))
	for i, f := range *testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeStringField(*f)
		res[i] = &sanitizedString
	}

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, &res)
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

	var res StringType
	res = StringType(testSanitizer.SanitizeStringField(string(testField)))

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
}

func TestEmbeddedStruct(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	type EmbeddedTestStruct struct {
		TestField string `sanitize:"test"`
	}

	testField := "test_field"
	testStruct := struct {
		EmbeddedTestStruct
	}{
		EmbeddedTestStruct{TestField: testField},
	}

	res := testSanitizer.SanitizeStringField(testField)

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
}

func TestNestedStruct(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	type nestedTestStruct struct {
		TestField string `sanitize:"test"`
	}

	testField := "test_field"
	testStruct := struct {
		Nested nestedTestStruct
	}{
		nestedTestStruct{TestField: testField},
	}

	res := testSanitizer.SanitizeStringField(testField)

	err := SanitizeInput(&testStruct)
	require.NoError(t, err)

	assert.Equal(t, testStruct.Nested.TestField, res)
}

func TestLocalSanitizer(t *testing.T) {
	testSanitizer := TestFieldSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField string `sanitize:"test"`
	}{
		TestField: testField,
	}

	res := testSanitizer.SanitizeStringField(testField)

	err := SanitizeInputWithLocalSanitizer(&testStruct, testSanitizer)
	require.NoError(t, err)

	assert.Equal(t, testStruct.TestField, res)
}

func TestInvalidSanitizer(t *testing.T) {

	testField := "test_field"
	testStruct := struct {
		TestField string `sanitize:"invalid"`
	}{
		TestField: testField,
	}

	err := SanitizeInput(&testStruct)
	assert.Equal(t, err.Error(), "invalid sanitizer")
}

func TestMissingOrEmptyTag(t *testing.T) {

	HtmlSanitizer := NewHtmlSanitizer(map[string]HtmlSanitizeOptions{}, false)

	testField := "test_field"
	testStruct := struct {
		TestField string
	}{
		TestField: testField,
	}

	err := SanitizeInputWithLocalSanitizer(&testStruct, HtmlSanitizer)
	require.Equal(t, err.Error(), "received empty tag on field \"TestField\" this is not allowed unless allowEmptyTag is explicitly set")
}

func TestInvalidMapField(t *testing.T) {
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

	err := SanitizeInput(&testStruct)
	require.Errorf(t, err, "cannot sanitize map inputs as they are not addressable")
}
