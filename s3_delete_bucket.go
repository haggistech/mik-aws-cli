package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

// Deletes an S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.

func deleteBucket() {
	if len(os.Args) != 3 {
		exitErrorf("bucket name required\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[2]

	// Set Region and use creds from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	// Create S3 client
	svc := s3.New(sess)

	// Delete the S3 Bucket - It must be empty or else the call fails
	_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait until bucket is deleted....
	fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully deleted\n", bucket)
}
