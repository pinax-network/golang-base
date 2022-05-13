package sanitizer

import (
	"github.com/microcosm-cc/bluemonday"
	"html"
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
	defaultOptions SanitizeOptions
	exceptions     map[string]SanitizeOptions
}

func NewHtmlSanitizer(defaultOptions SanitizeOptions, exceptions map[string]SanitizeOptions) *HtmlSanitizer {
	return &HtmlSanitizer{
		defaultOptions: defaultOptions,
		exceptions:     exceptions,
	}
}

func (h *HtmlSanitizer) SanitizeString(fieldName, fieldValue string) (res string) {

	var options SanitizeOptions

	if _, isException := h.exceptions[fieldName]; isException {
		options = h.exceptions[fieldName]
	} else {
		options = h.defaultOptions
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
