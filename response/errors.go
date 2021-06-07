package response

import "net/http"

var (
	Unauthorized        = NewApiError(http.StatusUnauthorized, UNAUTHORIZED)
	Forbidden           = NewApiError(http.StatusForbidden, FORBIDDEN)
	RouteNotFound       = NewApiError(http.StatusNotFound, NOT_FOUND_ROUTE)
	MethodNotAllowed    = NewApiError(http.StatusMethodNotAllowed, METHOD_NOT_ALLOWED)
	InternalServerError = NewApiError(http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
)

const (
	BAD_REQUEST_FILE_INPUT            = "bad_file_input"
	BAD_REQUEST_HEADER_MISSING        = "missing_required_header"
	BAD_REQUEST_HEADER                = "bad_header"
	BAD_REQUEST_JSON_INPUT            = "bad_json_input"
	BAD_REQUEST_QUERY_INPUT           = "bad_query_input"
	BAD_REQUEST_URI_INPUT             = "bad_uri_input"
	BAD_REQUEST_REGISTRATION_REQUIRED = "bad_request_registration_required"

	UNAUTHORIZED = "unauthorized"
	FORBIDDEN    = "forbidden"

	NOT_FOUND_ABUSE_REPORT      = "abuse_report_not_found"
	NOT_FOUND_FILE              = "file_not_found"
	NOT_FOUND_GRANT             = "grant_not_found"
	NOT_FOUND_GRANT_CATEGORY    = "grant_category_not_found"
	NOT_FOUND_GRANT_TRANSLATION = "grant_translation_not_found"
	NOT_FOUND_LANGUAGE          = "language_not_found"
	NOT_FOUND_REGION            = "region_not_found"
	NOT_FOUND_RESOURCE          = "resource_not_found"
	NOT_FOUND_ROUTE             = "route_not_found"
	NOT_FOUND_USER              = "user_not_found"

	METHOD_NOT_ALLOWED = "method_not_allowed"

	CONFLICT_EMAIL_EXISTS             = "email_already_exists"
	CONFLICT_GRANT_NAME_EXISTS        = "grant_name_already_exists"
	CONFLICT_GRANT_TRANSLATION_EXISTS = "grant_translation_already_exists"
	CONFLICT_USERNAME_EXISTS          = "username_already_exists"
	CONFLICT_PROFILE_EXISTS           = "user_profile_already_exists"

	INTERNAL_SERVER_ERROR = "internal_server_error"
)

type ApiErrorResponse struct {
	Errors []*ApiError `json:"errors"`
}

type ApiError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Detail string `json:"detail,omitempty"`
}

func (a *ApiError) Error() string {
	return a.Code
}

func (a *ApiError) Is(apiError *ApiError) bool {
	return a.Status == apiError.Status && a.Code == apiError.Code
}

func NewApiError(status int, code string) *ApiError {
	return &ApiError{
		Status: status,
		Code:   code,
	}
}

func NewApiErrorDetail(status int, code string, detail string) *ApiError {
	res := NewApiError(status, code)
	res.Detail = detail
	return res
}

func NewApiErrorBadRequestDetail(code string, detail string) *ApiError {
	res := NewApiError(http.StatusBadRequest, code)
	res.Detail = detail
	return res
}

func NewApiErrorBadRequest(code string) *ApiError {
	return NewApiErrorBadRequestDetail(code, "")
}

func NewApiErrorNotFoundDetail(code string, detail string) *ApiError {
	res := NewApiError(http.StatusNotFound, code)
	res.Detail = detail
	return res
}

func NewApiErrorNotFound(code string) *ApiError {
	return NewApiErrorNotFoundDetail(code, "")
}

func NewApiErrorConflictDetail(code string, detail string) *ApiError {
	res := NewApiError(http.StatusConflict, code)
	res.Detail = detail
	return res
}

func NewApiErrorConflict(code string) *ApiError {
	return NewApiErrorConflictDetail(code, "")
}
