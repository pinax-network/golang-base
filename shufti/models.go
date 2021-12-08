package shufti

type VerificationRequest struct {
	Reference         string               `json:"reference"`
	CallbackUrl       string               `json:"callback_url"`
	RedirectUrl       string               `json:"redirect_url"`
	Email             string               `json:"email"`
	Country           string               `json:"country"`
	VerificationMode  string               `json:"verification_mode"`
	Language          string               `json:"language,omitempty"`
	AllowOffline      string               `json:"allow_offline"`
	AllowOnline       string               `json:"allow_online"`
	ShowPrivacyPolicy string               `json:"show_privacy_policy"`
	ShowResults       string               `json:"show_results"`
	ShowConsent       string               `json:"show_consent"`
	ShowFeedbackForm  string               `json:"show_feedback_form"`
	Ttl               int                  `json:"ttl"`
	Document          DocumentVerification `json:"document"`
	BackgroundChecks  string               `json:"background_checks"`
}

type DocumentVerification struct {
	Proof             string   `json:"proof"`
	SupportedTypes    []string `json:"supported_types"`
	Name              string   `json:"name"`
	Dob               string   `json:"dob"`
	IssueDate         string   `json:"issue_date"`
	ExpiryDate        string   `json:"expiry_date"`
	DocumentNumber    string   `json:"document_number"`
	FetchEnhancedData string   `json:"fetch_enhanced_data"`
	Gender            string   `json:"gender"`
	PlaceOfIssue      string   `json:"place_of_issue"`
}

type VerificationResponse struct {
	Reference       string `json:"reference"`
	Event           string `json:"event"`
	Error           string `json:"error"`
	VerificationUrl string `json:"verification_url"`
}

type CallbackResponse struct {
	Reference      string   `json:"reference"`
	Event          string   `json:"event"`
	DeclinedReason string   `json:"declined_reason,omitempty"`
	DeclinedCodes  []string `json:"declined_codes,omitempty"`
}

func getDefaultDocumentVerificationRequest() VerificationRequest {
	return VerificationRequest{
		VerificationMode:  "any",
		AllowOffline:      "1",
		AllowOnline:       "1",
		ShowPrivacyPolicy: "1",
		ShowResults:       "1",
		ShowConsent:       "1",
		ShowFeedbackForm:  "0",
		Document: DocumentVerification{
			Proof:             "",
			SupportedTypes:    []string{"id_card", "driving_license", "passport"},
			Name:              "",
			Dob:               "",
			IssueDate:         "",
			ExpiryDate:        "",
			DocumentNumber:    "",
			FetchEnhancedData: "",
			Gender:            "",
			PlaceOfIssue:      "",
		},
		BackgroundChecks: "",
	}
}
