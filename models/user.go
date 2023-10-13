package base_models

import "time"

type User struct {
	ID            int       `json:"-"`                              // ID represents the user id within the local system (for example, MySQL primary keys)
	GUID          string    `json:"eosn_id" extensions:"x-order=0"` // GUID represents the unique global user id (such as the eosn_id or uuid)
	Email         string    `json:"email" extensions:"x-order=1"`
	EmailVerified bool      `json:"email_verified" extensions:"x-order=2"`
	Permissions   []string  `json:"-"`
	CreatedAt     time.Time `json:"created_at" extensions:"x-order=10"`
	UpdatedAt     time.Time `json:"udpated_at" extensions:"x-order=11"`
}
