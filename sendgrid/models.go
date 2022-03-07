package sendgrid

// Requests

type AddContactsRequest struct {
	ListIds  []string  `json:"list_ids"`
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	Email        string      `json:"email"`
	CustomFields interface{} `json:"custom_fields,omitempty"`
}

// Responses

type ImportContactsResponse struct {
	JobId string `json:"job_id"`
}
