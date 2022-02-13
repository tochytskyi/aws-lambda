package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadObject(bucket string, item string, fileName string, sess *session.Session) error {
	// open file to save contents to
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Unable to open file %q, %v", err)
		return err
	}
	defer file.Close()

	// Download from s3
	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		},
	)
	if err != nil {
		fmt.Printf("Unable to download item %q, %v", item, err)
		return err
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return nil
}
