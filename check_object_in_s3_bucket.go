package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
//	"github.com/aws/aws-sdk-go/aws/awserr"
)

func main() {
	s3region := os.Args[1]
	s3bucket := os.Args[2]
	s3object := os.Args[3]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3region)},
	)
	svc := s3.New(sess)
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(s3object),
	}

	_, err = svc.HeadObject(input)
	if err == nil {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
	os.Exit(0)
}
