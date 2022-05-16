package base_input

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/volatiletech/null/v8"
	"html"
	"strings"
)

var (
	strictSanitizer = bluemonday.StrictPolicy()
	htmlSanitizer   = bluemonday.UGCPolicy()
)

type SanitizeOptions struct {
	Strict           bool
	UnescapeString   bool
	CleanWhitespaces bool
}

func GetDefaultStrictOptions() SanitizeOptions {
	return SanitizeOptions{
		Strict:           true,
		UnescapeString:   true,
		CleanWhitespaces: true,
	}
}

func GetDefaultHtmlOptions() SanitizeOptions {
	return SanitizeOptions{
		Strict:           false,
		UnescapeString:   false,
		CleanWhitespaces: true,
	}
}

// SanitizeString will apply a html sanitization on the given string. If strict is set to true all html tags will be
// removed (bluemonday.StrictPolicy() will be applied). Otherwise, it will only remove html tags and attributes that are
// deemed dangerous (bluemonday.UGCPolicy() will be applied).
func SanitizeString(input string, options SanitizeOptions) string {

	res := input

	if options.Strict {
		res = strictSanitizer.Sanitize(res)
	} else {
		res = htmlSanitizer.Sanitize(res)
	}

	if options.UnescapeString {
		res = html.UnescapeString(res)
	}

	if options.CleanWhitespaces {
		res = strings.TrimSpace(res)
	}

	return res
}

// SanitizeNullString provides the functionality of SanitizeString for the null.String datatype. If the input
// is not valid, it will return the input variable. Otherwise it returns the result of SanitizeString as null.String
func SanitizeNullString(input null.String, options SanitizeOptions) null.String {

	if !input.Valid {
		return input
	}

	return null.StringFrom(SanitizeString(input.String, options))
}

// SanitizeStringSlice provides the functionality of SanitizeString for the string slices.
func SanitizeStringSlice(input []string, options SanitizeOptions) []string {

	res := make([]string, len(input))

	for i, str := range input {
		res[i] = SanitizeString(str, options)
	}

	return res
}
