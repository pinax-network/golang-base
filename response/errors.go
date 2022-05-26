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
	BAD_REQUEST_FILE_INPUT            = "bad_file_input"                    // request required a file but is not
	BAD_REQUEST_FILE_NOT_IMAGE        = "bad_file_not_an_image"             // request required an image but is not
	BAD_REQUEST_FILE_TOO_LARGE        = "bad_file_too_large"                // given file exceeds the filesize limit
	BAD_REQUEST_HEADER                = "bad_header"                        // invalid or malformed header given
	BAD_REQUEST_HEADER_MISSING        = "missing_required_header"           // request is missing a header
	BAD_REQUEST_JSON_INPUT            = "bad_json_input"                    // given json input is missing or malformed
	BAD_REQUEST_QUERY_INPUT           = "bad_query_input"                   // given query input is missing or malformed
	BAD_REQUEST_REGISTRATION_REQUIRED = "bad_request_registration_required" // user profile required for accessing this endpoint
	BAD_REQUEST_URI_INPUT             = "bad_uri_input"                     // given uri input is missing or malformed

	UNAUTHORIZED = "unauthorized" // invalid authorization information given

	FORBIDDEN = "forbidden" // not allowed to access this endpoint

	NOT_FOUND_ABUSE_REPORT            = "abuse_report_not_found"            // the requested abuse report was not found
	NOT_FOUND_BOUNTY                  = "bounty_not_found"                  // the requested bounty was not found
	NOT_FOUND_BOUNTY_CATEGORY         = "bounty_category_not_found"         // the requested bounty category was not found
	NOT_FOUND_BOUNTY_SKILL            = "bounty_skill_not_found"            // the requested bounty skill was not found
	NOT_FOUND_BOUNTY_EXPERIENCE_LEVEL = "bounty_experience_level_not_found" // the requested bounty experience level was not found
	NOT_FOUND_FILE                    = "file_not_found"                    // the requested file was not found
	NOT_FOUND_FUNDING_ADDRESS         = "funding_address_not_found"         // the eos account for the funding address could not be found on chain
	NOT_FOUND_GRANT                   = "grant_not_found"                   // the requested grant was not found
	NOT_FOUND_GRANT_CATEGORY          = "grant_category_not_found"          // the requested grant category was not found
	NOT_FOUND_GRANT_TRANSLATION       = "grant_translation_not_found"       // the requested grant translation was not found
	NOT_FOUND_KYC_REQUIRED            = "kyc_not_required"                  // kyc information has been requested but kyc is not required for the given user
	NOT_FOUND_LANGUAGE                = "language_not_found"                // the requested language was not found
	NOT_FOUND_LINKED_ACCOUNT          = "linked_account_not_found"          // the given account name is not linked to an eosn id
	NOT_FOUND_LINKED_SOCIAL           = "linked_social_not_found"           // the user tried to unlink a social that hasn't been linked before
	NOT_FOUND_MATCHING_PARTNER        = "matching_partner_not_found"        // the requested matching partner was not found
	NOT_FOUND_MATCHING_ROUND          = "matching_round_not_found"          // the requested matching round was not found
	NOT_FOUND_PORT_ACCOUNT            = "port_account_not_found"            // there has been no port verification found for the authenticated eosn_id
	NOT_FOUND_REGION                  = "region_not_found"                  // the requested region was not found
	NOT_FOUND_RESOURCE                = "resource_not_found"                // the requested resource was not found
	NOT_FOUND_ROUTE                   = "route_not_found"                   // the requested route was not found
	NOT_FOUND_SEASON                  = "season_not_found"                  // the requested season was not found
	NOT_FOUND_TOKEN                   = "token_not_found"                   // the requested token was not found
	NOT_FOUND_USER                    = "user_not_found"                    // the requested user was not found
	NOT_FOUND_USER_CHAIN_ACCOUNT      = "user_chain_account_not_found"      // there is no eosn chain account for the given eosn id
	NOT_FOUND_USER_PROFILE            = "user_profile_not_found"            // the requested user profile was not found

	METHOD_NOT_ALLOWED = "method_not_allowed" // http method is not allowed on this endpoint

	CONFLICT_BOUNTY_NAME_EXISTS            = "bounty_name_already_exists"        // the given bounty name has already been used
	CONFLICT_CANNOT_UNLINK_MAIN_SOCIAL     = "cannot_unlink_main_social"         // the user tried to unlink the social provider he has used to create the account
	CONFLICT_EMAIL_EXISTS                  = "email_already_exists"              // the given email already exists in the database
	CONFLICT_GRANT_ALREADY_APPLIED         = "grant_already_applied"             // the given grant has already applied to this matching round
	CONFLICT_GRANT_NAME_EXISTS             = "grant_name_already_exists"         // the given grant name already exists in the database
	CONFLICT_GRANT_NOT_PUBLISHED           = "grant_not_published"               // the given grant is not published yet
	CONFLICT_KYC_ALREADY_APPROVED          = "kyc_already_approved"              // Kyc verification url has been requested but the user has already been approved kyc
	CONFLICT_KYC_ALREADY_DECLINED          = "kyc_already_declined"              // Kyc verification url has been requested but the user has already been declined kyc
	CONFLICT_KYC_ALREADY_RECEIVED          = "kyc_already_received"              // Kyc verification url has been requested but the user has gone though the kyc process (but no result available yet)
	CONFLICT_LINKED_ACCOUNT_EXISTS         = "account_already_linked"            // the given eos account is already linked to an eosn id
	CONFLICT_LINKED_ACCOUNT_VERIFIED       = "linked_account_verified"           // the given linked account name cannot be deleted as it's already verified on chain (needs unlink transaction)
	CONFLICT_MAX_ACCOUNTS_LINKED           = "max_accounts_linked"               // the maximum amount of accounts has already been linked
	CONFLICT_MAX_PENDING_GRANTS            = "max_pending_grants_reached"        // the maximum of pending grants is reached for this user
	CONFLICT_MAX_TOTAL_GRANTS              = "max_total_grants_reached"          // the maximum amount of grants is reached for this user
	CONFLICT_MIN_REWARD_NOT_REACHED        = "min_reward_amount_not_reached"     // the minimum reward amount for a bounty has not been reached
	CONFLICT_NOT_IN_ACTIVE_MATCHING_ROUND  = "not_in_active_matching_round"      // the grant is not in an active matching round
	CONFLICT_PORT_VERIFICATION_EXPIRED     = "port_verification_expired"         // the port verification has already been expired
	CONFLICT_PORT_ATTESTATION_INSUFFICIENT = "port_attestation_insufficient"     // a port verification has been found but the attestation level is insufficient
	CONFLICT_PROFILE_EXISTS                = "user_profile_already_exists"       // a user profile already exists for this user
	CONFLICT_ROUND_ID_EXISTS               = "matching_round_id_already_exists"  // a matching round with the given id already exists in the database
	CONFLICT_SEASON_ID_EXISTS              = "season_id_already_exists"          // a season with the given id already exists in the database
	CONFLICT_SUBMISSIONS_NOT_OPEN          = "submissions_not_open"              // a grant has tried to apply to a matching round outside the submission range
	CONFLICT_SOCIAL_ALREADY_CONNECTED      = "social_provider_already_connected" // the user has already connected his account with the given social provider
	CONFLICT_USERNAME_EXISTS               = "username_already_exists"           // the given username already exists in the database
	CONFLICT_USER_EXISTS                   = "user_already_exists"               // a user with the given eosn id already exists in the database

	INTERNAL_SERVER_ERROR = "internal_server_error" // an unknown error occurred on the backend
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
