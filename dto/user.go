package dto

type User struct {
	ID          int      `json:"-"`
	EosnID      string   `json:"-" extensions:"x-order=0"`
	Permissions []string `json:"-"`
}
