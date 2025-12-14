package cloudstorage

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CloudFlyConfig struct {
	client     *s3.Client
	bucketName string
	endpoint   string
}

// NewCloudFlyConfig khởi tạo cấu hình kết nối đến CloudFly

// lưu trữ đám mây dựa trên AWS S3
func NewCloudFlyConfig(endpoint, accessKey, secretKey, bucketName string) (*CloudFlyConfig, error) {
	// Implementation for initializing CloudFlyConfig

	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           endpoint,
					SigningRegion: "auto",
				}, nil
			}),
		Region: "auto",
	}
	client := s3.NewFromConfig(cfg)
	return &CloudFlyConfig{
		client:     client,
		bucketName: bucketName,
		endpoint:   endpoint,
	}, nil
}
func (c *CloudFlyConfig) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	// đọc nội dung file
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(file); err != nil {
		return "", err
	}

	// tạo tên file duy nhất
	key := fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), filepath.Base(fileHeader.Filename))

	_, err := c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    "public-read", // cho phép truy cập công khai (optional)
	})
	if err != nil {
		return "", err
	}

	// tạo URL public
	url := fmt.Sprintf("%s/%s/%s", c.endpoint, c.bucketName, key)
	return url, nil
}
