package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
	"os"
)

func CopyImage(bucket string, item string, copyToBucket string, awsS3Session *s3.S3) {
	source := bucket + "/" + item

	_, err := awsS3Session.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(copyToBucket),
		CopySource: aws.String(url.PathEscape(source)),
		Key:        aws.String(item),
	})

	if err != nil {
		exitErrorf("Unable to copy item from bucket %q to bucket %q, %v", bucket, copyToBucket, err)
	}

	// Wait to see if the item got copied
	err = awsS3Session.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(copyToBucket),
		Key:    aws.String(item),
	})

	if err != nil {
		exitErrorf("Error occurred while waiting for item %q to be copied to bucket %q, %v", bucket, item, copyToBucket, err)
	}

	fmt.Printf("Item %q successfully copied from bucket %q to bucket %q\n", item, bucket, copyToBucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
