package dto

type User struct {
	EosnID      string   `json:"-" extensions:"x-order=0"`
	Permissions []string `json:"-"`
}
