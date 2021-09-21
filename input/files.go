package base_input

import "mime/multipart"

type FileInput struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
