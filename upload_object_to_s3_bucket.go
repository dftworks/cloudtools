package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"log"
	"bytes"
	"net/http"
	//	"github.com/aws/aws-sdk-go/aws/awserr"
)

func main() {
	s3region := os.Args[2]
	s3bucket := os.Args[3]
	s3object := os.Args[4]
	localfile := os.Args[1]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3region)},
	)
	svc := s3.New(sess)

	upFile, err := os.Open(localfile)
	if err != nil {
		log.Fatal(err)
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	input := &s3.PutObjectInput{
		Bucket:             aws.String(s3bucket),
		Key:                aws.String(s3object),
		Body:               bytes.NewReader(fileBuffer),
		ContentLength:      aws.Int64(fileSize),
		ContentType:        aws.String(http.DetectContentType(fileBuffer)),
	}

	_, err = svc.PutObject(input)
	if err == nil {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
	os.Exit(0)
}
