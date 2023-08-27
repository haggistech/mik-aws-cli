package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Printf("Usage: go run s3.go <the bucket name> <the AWS Region to use>\n" +
			"Example: go run s3.go my-test-bucket us-east-2\n")
		os.Exit(1)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("eu-west-1"),
	})

	arg := os.Args[1]

	switch {
	case arg == "s3cb":
		listMyBuckets(svc)
		createBucket()
	case arg == "s3db":
		deleteBucket()
		listMyBuckets(svc)
	case arg == "s3up":
		uploadObject()
		listMyBuckets(svc)
	case arg == "s3ls":
		listObjects()
	case arg == "sns-list":
		listTopics()
	case arg == "sns-create":
		snsCreate()
	case arg == "sns-subscribe":
		topicSubscribe()
	}

}

// List all of your available buckets in this AWS Region.
func listMyBuckets(svc *s3.S3) {
	result, err := svc.ListBuckets(nil)

	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Current Buckets:\n")

	for _, b := range result.Buckets {
		fmt.Printf(aws.StringValue(b.Name) + "\n")
	}

	fmt.Printf("\n")
}

// If there's an error, display it.
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
