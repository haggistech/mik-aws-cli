package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"

    "fmt"
    "os"
)

func topicSubscribe() {
    email := os.Args[2]
    topic := os.Args[3]


    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file. (~/.aws/credentials).
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := sns.New(sess)

    result, err := svc.Subscribe(&sns.SubscribeInput{
        Endpoint:              aws.String(email),
        Protocol:              aws.String("email"),
        ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
        TopicArn:              aws.String(topic),
    })
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(*result.SubscriptionArn)
}