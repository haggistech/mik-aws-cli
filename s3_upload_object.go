package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func uploadObject() {
	if len(os.Args) != 4 {
		exitErrorf("bucket and file name required\nUsage: %s bucket_name filename",
			os.Args[0])
	}

	bucket := os.Args[2]
	filename := os.Args[3]

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", filename, err)
	}

	defer file.Close()

	// Set Region and use creds from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	// Setup the S3 Upload Manager
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(sess)

	// Upload the file's body to S3 bucket key being the same as the filename.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		Key: aws.String(filename),

		Body: file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}
