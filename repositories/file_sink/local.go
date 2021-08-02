package file_sink

import (
	"context"
	"fmt"
	base_repositories "github.com/eosnationftw/eosn-base-api/repositories"
	"golang.org/x/sys/unix"
	"net/url"
	"os"
	"path"
)

type LocalFileRepository struct {
	baseUrl string
	subDirs map[base_repositories.FileType]string
}

func NewLocalFileRepository(subDirs map[base_repositories.FileType]string) (*LocalFileRepository, error) {

	// check if we have write access on all given sub dirs
	for _, dir := range subDirs {
		_, err := getStaticFileDir(dir)
		if err != nil {
			return nil, err
		}
	}

	baseUrl := os.Getenv("STATIC_BASE_URL")
	if baseUrl == "" {
		return nil, fmt.Errorf("env STATIC_BASE_URL is not set")
	}

	return &LocalFileRepository{
		baseUrl: baseUrl,
		subDirs: subDirs,
	}, nil
}

func (f *LocalFileRepository) Init() error {
	return nil
}

func (f *LocalFileRepository) FileExists(ctx context.Context, fileUuid string, fileType base_repositories.FileType) bool {

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

func (f *LocalFileRepository) UploadFile(ctx context.Context, tmpFile, fileUuid string, fileType base_repositories.FileType) {

	targetFile := f.getStaticFileName(fileUuid, fileType)

	err := os.Rename(tmpFile, targetFile)
	if err != nil {
		panic(fmt.Sprintf("failed to move temp file from temp dir ('%s') to static file dir ('%s')", tmpFile, targetFile))
	}
}

func (f *LocalFileRepository) GetFileUrl(ctx context.Context, fileUuid string, fileType base_repositories.FileType) string {

	subDir, ok := f.subDirs[fileType]
	if !ok {
		panic("no subdir initialized for given file type: " + fileType)
	}

	return mustJoinUrl(f.baseUrl, path.Join(subDir, fileUuid))
}

func getStaticFileDir(subDir string) (staticDir string, err error) {
	staticDir = path.Join(os.Getenv("STATIC_FILE_DIR"), subDir)

	if err = writeable(staticDir); err != nil {
		err = fmt.Errorf("static file dir ('%s') not writable: '%e'", staticDir, err)
	}

	return
}

func (f *LocalFileRepository) getStaticFileName(fileUuid string, fileType base_repositories.FileType) string {

	subDir, ok := f.subDirs[fileType]
	if !ok {
		panic("no subdir initialized for given file type: " + fileType)
	}

	return path.Join(subDir, fileUuid)
}

func mustJoinUrl(baseurl, urlPath string) string {
	u, err := url.Parse(baseurl)
	if err != nil {
		panic(fmt.Errorf("failed to parse base url for static file serving: '%s', error: %e", baseurl, err))
	}

	u.Path = path.Join(u.Path, urlPath)
	return u.String()
}

func writeable(dir string) error {
	return unix.Access(dir, unix.W_OK)
}
