package file_sink

import (
	"context"
	"fmt"
	base_repositories "github.com/pinax-network/golang-base/repositories"
	"golang.org/x/sys/unix"
	"net/url"
	"os"
	"path"
)

type LocalFileRepository struct {
	config   *LocalFileSinkConfig
	subDirs  map[base_repositories.FileType]string
	fileDirs map[base_repositories.FileType]string
}

// NewLocalFileRepository creates a new local file repository. The subDirs parameter specifies the subdirectories within
// the STATIC_FILE_DIR which is specified as env variable. So if STATIC_FILE_DIR is /var/www/static and the subdir for
// the USER_AVATAR base_repositories.FileType is avatars, then the static directory for avatars will be /var/www/static/avatars/
//
// The files will then be linked under STATIC_BASE_URL within the subDirs parameter. So if STATIC_BASE_URL is
// http://localhost:8080/static user avatars will be linked http://localhost:8080/static/avatars/
func NewLocalFileRepository(subDirs map[base_repositories.FileType]string, config *LocalFileSinkConfig) (*LocalFileRepository, error) {

	fileDirs := make(map[base_repositories.FileType]string)

	// replace sub dirs with static file dir path and check if it is writeable
	for key, dir := range subDirs {
		staticFileDir, err := getStaticFileDir(config.UploadDir, dir)
		if err != nil {
			return nil, err
		}
		fileDirs[key] = staticFileDir
	}

	return &LocalFileRepository{
		config:   config,
		subDirs:  subDirs,
		fileDirs: fileDirs,
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

	return mustJoinUrl(f.config.BaseUrl, path.Join(subDir, fileUuid))
}

func getStaticFileDir(baseDir, subDir string) (staticDir string, err error) {
	staticDir = path.Join(baseDir, subDir)

	if err = writeable(staticDir); err != nil {
		err = fmt.Errorf("static file dir ('%s') not writable: '%e'", staticDir, err)
	}

	return
}

func (f *LocalFileRepository) getStaticFileName(fileUuid string, fileType base_repositories.FileType) string {

	fileDir, ok := f.fileDirs[fileType]
	if !ok {
		panic("no subdir initialized for given file type: " + fileType)
	}

	return path.Join(fileDir, fileUuid)
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
