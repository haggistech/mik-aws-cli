package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"

    "fmt"
    "os"
)

func snsCreate() {
    if len(os.Args) < 2 {
        fmt.Println("You must supply a topic name")
        os.Exit(1)
    }

    // Set Region and use creds from ~/.aws/credentials
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := sns.New(sess)

    result, err := svc.CreateTopic(&sns.CreateTopicInput{
        Name: aws.String(os.Args[2]),
    })
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(*result.TopicArn)
}