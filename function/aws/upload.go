package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadObject(bucket string, filePath string, fileName string, sess *session.Session) error {
	// Open file to upload
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Unable to open file")
		return err
	}
	defer file.Close()

	// Upload to s3
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		fmt.Printf("failed to upload object")
		return err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	return nil
}
