package helper

import (
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
	"strings"
)

func UploadMimetypeContains(file *multipart.FileHeader, hasMime string) (contains bool, err error) {

	fileSrc, err := file.Open()
	if err != nil {
		return
	}

	fileHeader := make([]byte, 512)
	_, err = fileSrc.Read(fileHeader)
	if err != nil {
		return
	}

	mime := mimetype.Detect(fileHeader)
	if strings.Contains(mime.String(), "image") {
		contains = true
	}

	return
}
