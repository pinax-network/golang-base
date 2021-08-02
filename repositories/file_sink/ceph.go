package repositories

import (
	"context"
	"fmt"
	"github.com/eosnationftw/eosn-base-api/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"net/url"
	"os"
	"path"
)

const CEPH_USER_AVATAR_SUBDIR = "avatars"

type CephRepository struct {
	minioClient *minio.Client
	bucket      string
	baseUrl     string
}

func NewCephRepository() (*CephRepository, error) {

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
	}, nil
}

func (c *CephRepository) FileExists(ctx context.Context, fileUuid string, fileType FileType) bool {

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

func (c *CephRepository) UploadFile(ctx context.Context, tmpFile, fileUuid string, fileType FileType) {

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

func (c *CephRepository) GetFileUrl(ctx context.Context, fileUuid string, fileType FileType) string {

	switch fileType {
	case USER_AVATAR:
		return c.mustJoinUrl(c.baseUrl, path.Join(CEPH_USER_AVATAR_SUBDIR, fileUuid))
	default:
		panic("invalid file type given: " + fileType)
	}
}

func (c *CephRepository) getStaticFileName(fileUuid string, fileType FileType) string {

	var res string
	switch fileType {
	case USER_AVATAR:
		res = path.Join(CEPH_USER_AVATAR_SUBDIR, fileUuid)
		break
	default:
		panic("invalid file type given: " + fileType)
	}

	return res
}

func (c *CephRepository) mustJoinUrl(baseurl, urlPath string) string {
	u, err := url.Parse(baseurl)
	if err != nil {
		panic(fmt.Errorf("failed to parse base url for static file serving: '%s', error: %v", baseurl, err))
	}

	u.Path = path.Join(u.Path, urlPath)
	return u.String()
}
