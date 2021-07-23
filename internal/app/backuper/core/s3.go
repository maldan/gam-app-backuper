package core

import (
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/maldan/go-cmhp/cmhp_file"
)

var s3Client *s3.S3
var config Config

func StartS3() {
	cmhp_file.ReadJSON(DataDir+"/config.json", &config)

	newSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.SPACES_KEY, config.SPACES_SECRET, ""),
		Endpoint:    aws.String(config.SPACES_ENDPOINT),
		Region:      aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	s3Client = s3.New(newSession)
	log.Println("S3 INITED")
}

func PureS3Path(path string) string {
	// Url of endpoint
	url := strings.Replace(
		config.SPACES_ENDPOINT,
		"https://", "https://"+config.SPACES_BUCKET+".",
		1,
	)
	path = strings.Replace(path, url+"/", "", 1)
	return path
}

func List(path string) []string {
	list, _ := s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(config.SPACES_BUCKET),
		Prefix: aws.String(path),
	})
	out := make([]string, 0)
	for _, o := range list.Contents {
		out = append(out, *o.Key)
	}
	return out
}

func Upload(path string, data []byte, visibility string, contentType string) (string, error) {
	// Url of endpoint
	url := strings.Replace(
		config.SPACES_ENDPOINT,
		"https://", "https://"+config.SPACES_BUCKET+".",
		1,
	)

	// Remove https://...
	path = PureS3Path(path)

	object := s3.PutObjectInput{
		Bucket:      aws.String(config.SPACES_BUCKET),
		Key:         aws.String(path),
		Body:        bytes.NewReader(data),
		ACL:         aws.String(visibility),
		ContentType: aws.String(contentType),
	}

	_, err := s3Client.PutObject(&object)
	if err != nil {
		return "", err
	}

	return url + "/" + path, nil
}

func Download(path string) ([]byte, error) {
	// Remove https://...
	path = PureS3Path(path)

	// Download file
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.SPACES_BUCKET),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	// Read and decompress
	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
