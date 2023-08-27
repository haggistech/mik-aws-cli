package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func createBucket() {
	if len(os.Args) != 3 {
		exitErrorf("Bucket name missing!\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[2]

	// Set Region and use creds from ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	// Create S3 client
	svc := s3.New(sess)

	// Create the S3 Bucket
	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created.....
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully created\n", bucket)
}
