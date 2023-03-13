package sanitizer

// Sanitizer is the sanitizer that will be used whenever the SanitizeInput function is called.
//
// By default, this is set to a html sanitizer that allows to apply the 'strict' (strips away *all* html tags) and
// 'html' policy (strips away dangerous html tags).
//
// To set your own sanitizer just override this variable with your own before calling SanitizeInput.
var Sanitizer FieldSanitizer = NewHtmlSanitizer(map[string]SanitizeOptions{
	"strict": GetDefaultStrictOptions(),
	"html":   GetDefaultHtmlOptions(),
})
