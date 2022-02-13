package convert

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"io"
	"log"
	"strings"

	awsHelpers "aws-lambda/aws"
	"github.com/sunshineplan/imgconv"
)

var formatExtensions = map[imgconv.Format]string{
	imgconv.JPEG: "jpg",
	imgconv.PNG:  "png",
	imgconv.GIF:  "gif",
	imgconv.BMP:  "bmp",
}

func Convert(
	bucket string,
	item string,
	copyToBucket string,
	awsSession *session.Session,
	format imgconv.Format,
) {
	buffer := &aws.WriteAtBuffer{}

	err := awsHelpers.DownloadObject(bucket, item, buffer, awsSession)
	if err != nil {
		log.Fatalf("failed to download source image: %v", err)
	}

	src, err := imgconv.Decode(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	err = imgconv.Write(io.Discard, src, imgconv.FormatOption{Format: format})
	if err != nil {
		log.Fatalf("failed to convert to PNG: %v", err)
	}

	convertedFileName := strings.Split(item, ".")[0] + "." + formatExtensions[format]

	err = awsHelpers.UploadObject(copyToBucket, buffer, convertedFileName, awsSession)
	if err != nil {
		log.Fatalf("failed to upload converted PNG: %v", err)
	}
}
