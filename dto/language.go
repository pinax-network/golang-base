package dto

type Language struct {
	ID             int    `db:"id" json:"id"`
	Code           string `db:"code" json:"code"`
	Language       string `db:"language" json:"language"`
	LanguageNative string `db:"language_native" json:"language_native"`
}
