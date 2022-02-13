package aws

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadObject(bucket string, buffer *aws.WriteAtBuffer, fileName string, sess *session.Session) error {
	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer.Bytes()),
	})

	if err != nil {
		fmt.Printf("failed to upload object")
		return err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)

	return nil
}
