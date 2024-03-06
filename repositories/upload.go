package base_repositories

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	base_models "github.com/pinax-network/golang-base/models"
	"golang.org/x/sys/unix"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

type FileType string

const (
	USER_AVATAR FileType = "user_avatar"
)

type UploadRepository struct {
	tempUploadDir  string
	fileRepository StaticFileRepository
}

type StaticFileRepository interface {
	FileExists(ctx context.Context, fileUuid string, fileType FileType) bool
	UploadFile(ctx context.Context, tmpFile, fileUuid string, fileType FileType)
	GetFileUrl(ctx context.Context, fileUuid string, fileType FileType) string
}

func NewUploadRepository(fileRepository StaticFileRepository, config *UploadRepositoryConfig) (*UploadRepository, error) {

	if err := writeable(config.TempUploadDir); err != nil {
		err = fmt.Errorf("temp upload dir ('%s') not writable: '%e'", config.TempUploadDir, err)
	}

	return &UploadRepository{
		tempUploadDir:  config.TempUploadDir,
		fileRepository: fileRepository,
	}, nil
}

func (u *UploadRepository) SaveTempFile(c *gin.Context, file *multipart.FileHeader) (fileName *base_models.UploadedFile) {

	extension := filepath.Ext(file.Filename)
	fileNameUuid := uuid.New().String() + extension

	err := c.SaveUploadedFile(file, path.Join(u.tempUploadDir, fileNameUuid))
	if err != nil {
		panic(fmt.Errorf("failed to save uploaded file to temp storage: %v", err))
	}

	fileName = &base_models.UploadedFile{Filename: fileNameUuid}

	return
}

func (s *UploadRepository) UploadFile(ctx context.Context, fileUuid string, fileType FileType) string {
	s.fileRepository.UploadFile(ctx, path.Join(s.tempUploadDir, fileUuid), fileUuid, fileType)
	return s.fileRepository.GetFileUrl(ctx, fileUuid, fileType)
}

func (s *UploadRepository) GetFileUrl(ctx context.Context, fileUuid string, fileType FileType) string {
	return s.fileRepository.GetFileUrl(ctx, fileUuid, fileType)
}

func (s *UploadRepository) MustExistsTemp(ctx context.Context, fileUuid string) bool {

	if fileUuid == "" {
		return false
	}

	if _, err := os.Stat(path.Join(s.tempUploadDir, fileUuid)); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(fmt.Sprintf("failed to check if file exists: '%s', error: %e", fileUuid, err))
	}
}

func (s *UploadRepository) MustExists(ctx context.Context, fileUuid string, fileType FileType) bool {
	return s.fileRepository.FileExists(ctx, fileUuid, fileType)
}

func writeable(dir string) error {
	return unix.Access(dir, unix.W_OK)
}
