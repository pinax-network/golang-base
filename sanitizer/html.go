package sanitizer

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/volatiletech/null/v8"
	"html"
	"reflect"
	"strings"
)

type SanitizeOptions struct {
	Policy           *bluemonday.Policy
	UnescapeString   bool
	CleanWhitespaces bool
}

func GetDefaultStrictOptions() SanitizeOptions {
	return SanitizeOptions{
		Policy:           bluemonday.StrictPolicy(),
		UnescapeString:   true,
		CleanWhitespaces: true,
	}
}

func GetDefaultHtmlOptions() SanitizeOptions {
	return SanitizeOptions{
		Policy:           bluemonday.UGCPolicy(),
		UnescapeString:   false,
		CleanWhitespaces: true,
	}
}

type HtmlSanitizer struct {
	sanitizer map[string]SanitizeOptions
}

func NewHtmlSanitizer(sanitizer map[string]SanitizeOptions) *HtmlSanitizer {
	return &HtmlSanitizer{
		sanitizer: sanitizer,
	}
}

func (h *HtmlSanitizer) SanitizeString(field reflect.StructField, fieldValue string) (res string, err error) {

	fieldTag := field.Tag.Get(TagName)

	// if we don't have a value for the sanitize tag on this field, we just return the raw value
	if field.Tag.Get(TagName) == "" {
		res = fieldValue
		return
	}

	options, ok := h.sanitizer[fieldTag]

	if !ok {
		err = fmt.Errorf("failed to sanitize field %q, with sanitizer %q: no such sanitizer initialized", field.Name, fieldTag)
		return
	}

	res = options.Policy.Sanitize(fieldValue)

	if options.UnescapeString {
		res = html.UnescapeString(res)
	}

	if options.CleanWhitespaces {
		res = strings.TrimSpace(res)
	}

	return
}

func (h *HtmlSanitizer) SanitizeNullString(field reflect.StructField, value null.String) (res null.String, err error) {

	res = value

	if !value.Valid {
		return
	}

	sanitizedString, err := h.SanitizeString(field, value.String)
	if err != nil {
		return
	}

	res = null.StringFrom(sanitizedString)
	return
}
