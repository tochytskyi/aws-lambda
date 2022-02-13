package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadObject(bucket string, item string, buffer *aws.WriteAtBuffer, sess *session.Session) error {
	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		},
	)

	if err != nil {
		fmt.Printf("Unable to download item %q, %v", item, err)
		return err
	}

	fmt.Println("Downloaded item with ", numBytes, "bytes")

	return nil
}
