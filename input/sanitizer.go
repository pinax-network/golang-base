package base_input

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/volatiletech/null/v8"
)

var (
	strictSanitizer = bluemonday.StrictPolicy()
	htmlSanitizer   = bluemonday.UGCPolicy()
)

func SanitizeString(input string, strict bool) string {
	if strict {
		return strictSanitizer.Sanitize(input)
	}
	return htmlSanitizer.Sanitize(input)
}

func SanitizeNullString(input null.String, strict bool) null.String {

	if !input.Valid {
		return null.StringFromPtr(nil)
	}

	if strict {
		return null.StringFrom(strictSanitizer.Sanitize(input.String))
	} else {
		return null.StringFrom(htmlSanitizer.Sanitize(input.String))
	}
}
