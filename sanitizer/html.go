package sanitizer

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/volatiletech/null/v8"
	"html"
	"reflect"
	"strings"
)

const TagExclude = "-"

// HtmlSanitizeOptions define the options on how to sanitize fields using bluemonday.Policy for removing html tags.
type HtmlSanitizeOptions struct {
	Policy           *bluemonday.Policy
	UnescapeString   bool // unescapes special characters
	CleanWhitespaces bool // removes leading and trailing whitespaces
}

func GetDefaultStrictOptions() HtmlSanitizeOptions {
	return HtmlSanitizeOptions{
		Policy:           bluemonday.StrictPolicy(),
		UnescapeString:   true,
		CleanWhitespaces: true,
	}
}

func GetDefaultHtmlOptions() HtmlSanitizeOptions {
	return HtmlSanitizeOptions{
		Policy:           bluemonday.UGCPolicy(),
		UnescapeString:   false,
		CleanWhitespaces: true,
	}
}

type HtmlSanitizer struct {
	sanitizer     map[string]HtmlSanitizeOptions
	allowEmptyTag bool
}

// NewHtmlSanitizer initializes a new HtmlSanitizer with the given HtmlSanitizeOptions mapped to a tag name. If
// allowEmptyTag is set to false the sanitizer will return an error in case a string or null.String field does not have
// a sanitize tag set (or the tag is empty). To explicitly allow fields not being sanitized use "-" for the tag.
//
// Example initialization:
//
//	sanitizer := NewHtmlSanitizer(map[string]HtmlSanitizeOptions{
//		"strict": GetDefaultStrictOptions(),
//		"html":   GetDefaultHtmlOptions(),
//	}, false)
func NewHtmlSanitizer(sanitizer map[string]HtmlSanitizeOptions, allowEmptyTag bool) *HtmlSanitizer {
	return &HtmlSanitizer{
		sanitizer:     sanitizer,
		allowEmptyTag: allowEmptyTag,
	}
}

func (h *HtmlSanitizer) SanitizeString(field reflect.StructField, fieldValue string) (res string, err error) {

	fieldTag := field.Tag.Get(TagName)

	// in case we received an empty tag and allowEmptyTag is not explicitly set we return an error
	if field.Tag.Get(TagName) == "" && !h.allowEmptyTag {
		return "", fmt.Errorf("received empty tag on field %q this is not allowed unless allowEmptyTag is explicitly set", field.Name)
	}

	// in case empty fields are allowed or we received an empty tag we just pass the raw value
	if (field.Tag.Get(TagName) == "" && h.allowEmptyTag) || field.Tag.Get(TagName) == TagExclude {
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
