package repositories

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
)

const USER_AVATAR_SUBDIR = "avatars"

type FileRepository struct {
	userAvatarDir string
	baseUrl       string
}

func NewFileRepository() (*FileRepository, error) {

	userAvatarDir, err := getStaticFileDir(USER_AVATAR_SUBDIR)
	if err != nil {
		return nil, err
	}

	baseUrl := os.Getenv("STATIC_BASE_URL")
	if baseUrl == "" {
		return nil, fmt.Errorf("env STATIC_BASE_URL is not set")
	}

	return &FileRepository{
		userAvatarDir: userAvatarDir,
		baseUrl:       baseUrl,
	}, nil
}

func (f *FileRepository) Init() error {
	return nil
}

func (f *FileRepository) FileExists(ctx context.Context, fileUuid string, fileType FileType) bool {

	if fileUuid == "" {
		return false
	}

	file := f.getStaticFileName(fileUuid, fileType)

	if _, err := os.Stat(file); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(fmt.Sprintf("failed to check if file exists: '%s', error: %e", fileUuid, err))
	}
}

func (f *FileRepository) UploadFile(ctx context.Context, tmpFile, fileUuid string, fileType FileType) {

	targetFile := f.getStaticFileName(fileUuid, fileType)

	err := os.Rename(tmpFile, targetFile)
	if err != nil {
		panic(fmt.Sprintf("failed to move temp file from temp dir ('%s') to static file dir ('%s')", tmpFile, targetFile))
	}
}

func (f *FileRepository) GetFileUrl(ctx context.Context, fileUuid string, fileType FileType) string {

	switch fileType {
	case USER_AVATAR:
		return mustJoinUrl(f.baseUrl, path.Join(USER_AVATAR_SUBDIR, fileUuid))
	default:
		panic("invalid file type given: " + fileType)
	}
}

func getStaticFileDir(subDir string) (staticDir string, err error) {
	staticDir = path.Join(os.Getenv("STATIC_FILE_DIR"), subDir)

	if err = writeable(staticDir); err != nil {
		err = fmt.Errorf("static file dir ('%s') not writable: '%e'", staticDir, err)
	}

	return
}

func (f *FileRepository) getStaticFileName(fileUuid string, fileType FileType) string {

	var res string
	switch fileType {
	case USER_AVATAR:
		res = path.Join(f.userAvatarDir, fileUuid)
		break
	default:
		panic("invalid file type given: " + fileType)
	}

	return res
}

func mustJoinUrl(baseurl, urlPath string) string {
	u, err := url.Parse(baseurl)
	if err != nil {
		panic(fmt.Errorf("failed to parse base url for static file serving: '%s', error: %e", baseurl, err))
	}

	u.Path = path.Join(u.Path, urlPath)
	return u.String()
}
