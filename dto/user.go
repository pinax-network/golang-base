package dto

import "time"

type User struct {
	ID          int       `json:"-"`
	EosnID      string    `json:"eosn_id" extensions:"x-order=0"`
	Permissions []string  `json:"-"`
	CreatedAt   time.Time `json:"created_at" extensions:"x-order=1"`
	UpdatedAt   time.Time `json:"udpated_at" extensions:"x-order=2"`
}
