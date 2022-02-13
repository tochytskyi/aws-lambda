package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"

	"aws-lambda/convert"
	"github.com/sunshineplan/imgconv"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

var region = "us-east-2"
var copyToBucket = "jpeg-converted-images"

var awsSession *session.Session

//var awsS3Session *s3.S3

func handler(ctx context.Context, s3Event events.S3Event) (Response, error) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		//awsHelpers.CopyImage(s3.Bucket.Name, s3.Object.Key, copyToBucket, awsSession)
		convert.Convert(s3.Bucket.Name, s3.Object.Key, copyToBucket, awsSession, imgconv.PNG)
		convert.Convert(s3.Bucket.Name, s3.Object.Key, copyToBucket, awsSession, imgconv.BMP)
		convert.Convert(s3.Bucket.Name, s3.Object.Key, copyToBucket, awsSession, imgconv.GIF)
	}

	fmt.Printf("%d files converted", len(s3Event.Records))

	return Response{
		Message: fmt.Sprintf("%d files converted", len(s3Event.Records)),
		Ok:      true,
	}, nil
}

func main() {
	initSession()
	lambda.Start(handler)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func initSession() {
	sessionTmp, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AccessKeyID"),
			os.Getenv("AccessKeySecret"),
			"",
		),
	},
	)

	if err != nil {
		exitErrorf("Unable to init aws session")
	}

	awsSession = sessionTmp

	//awsS3Session = s3.New(awsSession)
}
