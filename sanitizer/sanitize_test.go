package sanitizer

import (
	"github.com/bmizerany/assert"
	"testing"
)

type TestSanitizer struct{}

func (t TestSanitizer) SanitizeString(fieldName, fieldValue string) string {
	return fieldValue + "_sanitized"
}

func TestStringField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField string
	}{
		TestField: testField,
	}

	res := SanitizeInput(testStruct, testSanitizer)
	testStruct.TestField = testSanitizer.SanitizeString("", testField)
	assert.Equal(t, testStruct, res)
}

func TestStringPtrField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testField := "test_field"
	testStruct := struct {
		TestField *string
	}{
		TestField: &testField,
	}

	res := SanitizeInput(testStruct, testSanitizer)
	sanitizedField := testSanitizer.SanitizeString("", testField)
	testStruct.TestField = &sanitizedField
	assert.Equal(t, testStruct, res)
}

func TestStringSliceField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testStruct := struct {
		TestField []string
	}{
		TestField: []string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res := SanitizeInput(testStruct, testSanitizer)

	for i, f := range testStruct.TestField {
		testStruct.TestField[i] = testSanitizer.SanitizeString("", f)
	}

	assert.Equal(t, testStruct, res)
}

func TestStringPtrSliceField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestField []*string
	}{
		TestField: []*string{&testString1, &testString2, &testString3},
	}

	res := SanitizeInput(testStruct, testSanitizer)

	for i, f := range testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeString("", *f)
		testStruct.TestField[i] = &sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

func TestStringSlicePtrField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testStruct := struct {
		TestField *[]string
	}{
		TestField: &[]string{"test_entry_1", "test_entry_2", "test_entry_3"},
	}

	res := SanitizeInput(testStruct, testSanitizer)

	for i, f := range *testStruct.TestField {
		(*testStruct.TestField)[i] = testSanitizer.SanitizeString("", f)
	}

	assert.Equal(t, testStruct, res)
}

func TestStringPtrSlicePtrField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestField *[]*string
	}{
		TestField: &[]*string{&testString1, &testString2, &testString3},
	}

	res := SanitizeInput(testStruct, testSanitizer)

	for i, f := range *testStruct.TestField {
		sanitizedString := testSanitizer.SanitizeString("", *f)
		(*testStruct.TestField)[i] = &sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

func TestMapField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	testString1 := "test_entry_1"
	testString2 := "test_entry_2"
	testString3 := "test_entry_3"

	testStruct := struct {
		TestMap map[string]string
	}{
		TestMap: map[string]string{
			"field1": testString1,
			"field2": testString2,
			"field3": testString3,
		},
	}

	res := SanitizeInput(testStruct, testSanitizer)

	for i, f := range testStruct.TestMap {
		sanitizedString := testSanitizer.SanitizeString("", f)
		testStruct.TestMap[i] = sanitizedString
	}

	assert.Equal(t, testStruct, res)
}

type StringType string

func TestStringTypeField(t *testing.T) {
	testSanitizer := TestSanitizer{}

	var testField StringType
	testField = "test_field"
	testStruct := struct {
		TestField StringType
	}{
		TestField: testField,
	}

	res := SanitizeInput(testStruct, testSanitizer)
	testStruct.TestField = StringType(testSanitizer.SanitizeString("", string(testField)))
	assert.Equal(t, testStruct, res)
}
