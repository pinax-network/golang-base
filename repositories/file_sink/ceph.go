package file_sink

import (
	"context"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/eosnationftw/eosn-base-api/repositories"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"net/url"
	"os"
	"path"
)

type CephRepository struct {
	minioClient *minio.Client
	bucket      string
	baseUrl     string
	subDirs     map[base_repositories.FileType]string
}

func NewCephRepository(subDirs map[base_repositories.FileType]string) (*CephRepository, error) {

	if os.Getenv("CEPH_HOST") == "" || os.Getenv("CEPH_ACCESS_KEY") == "" ||
		os.Getenv("CEPH_SECRET") == "" || os.Getenv("CEPH_SSL") == "" || os.Getenv("CEPH_BUCKET") == "" {

		return nil, fmt.Errorf("missing env variables, requires CEPH_HOST, CEPH_ACCESS_KEY, CEPH_SECRET, CEPH_SSL and CEPH_BUCKET")
	}

	minioClient, err := minio.New(os.Getenv("CEPH_HOST"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("CEPH_ACCESS_KEY"), os.Getenv("CEPH_SECRET"), ""),
		Secure: os.Getenv("CEPH_SSL") == "true",
	})
	if err != nil {
		return nil, err
	}

	return &CephRepository{
		minioClient: minioClient,
		bucket:      os.Getenv("CEPH_BUCKET"),
		baseUrl:     os.Getenv("STATIC_BASE_URL"),
		subDirs:     subDirs,
	}, nil
}

func (c *CephRepository) FileExists(ctx context.Context, fileUuid string, fileType base_repositories.FileType) bool {

	if fileUuid == "" {
		return false
	}

	file := c.getStaticFileName(fileUuid, fileType)
	_, err := c.minioClient.StatObject(ctx, c.bucket, file, minio.StatObjectOptions{})

	if err != nil {
		minioErr, ok := err.(minio.ErrorResponse)
		if !ok || minioErr.Code != "NoSuchKey" {
			log.Panic("failed to check if file exists", zap.Error(err))
		} else {
			return false
		}
	}

	return true
}

func (c *CephRepository) UploadFile(ctx context.Context, tmpFile, fileUuid string, fileType base_repositories.FileType) {

	targetFile := c.getStaticFileName(fileUuid, fileType)

	metaData := map[string]string{
		"X-Amz-Acl": "public-read",
	}

	_, err := c.minioClient.FPutObject(ctx, c.bucket, targetFile, tmpFile, minio.PutObjectOptions{
		UserMetadata: metaData,
	})
	if err != nil {
		log.Panic("failed to upload temp file to ceph", zap.Error(err), zap.String("tmp_file", tmpFile), zap.String("target_file", targetFile))
	}
}

func (c *CephRepository) GetFileUrl(ctx context.Context, fileUuid string, fileType base_repositories.FileType) string {

	subDir, ok := c.subDirs[fileType]
	if !ok {
		panic("no subdir initialized for given file type: " + fileType)
	}

	return c.mustJoinUrl(c.baseUrl, path.Join(subDir, fileUuid))
}

func (c *CephRepository) getStaticFileName(fileUuid string, fileType base_repositories.FileType) string {

	subDir, ok := c.subDirs[fileType]
	if !ok {
		panic("no subdir initialized for given file type: " + fileType)
	}

	return path.Join(subDir, fileUuid)
}

func (c *CephRepository) mustJoinUrl(baseurl, urlPath string) string {
	u, err := url.Parse(baseurl)
	if err != nil {
		panic(fmt.Errorf("failed to parse base url for static file serving: '%s', error: %v", baseurl, err))
	}

	u.Path = path.Join(u.Path, urlPath)
	return u.String()
}
