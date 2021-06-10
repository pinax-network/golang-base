package dto

import "time"

type User struct {
	ID          int       `json:"-"`
	EosnID      string    `json:"-" extensions:"x-order=0"`
	Permissions []string  `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
