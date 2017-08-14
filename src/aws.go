package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

func InitAws() *session.Session {
	sess, err := session.NewSession(&aws.Config{
    Region:      aws.String("ap-northeast-1"),
    Credentials: credentials.NewStaticCredentials(GetAwsInfo("access_key"), GetAwsInfo("secret_key"), ""),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return sess
}

func ShowList() {
	client := s3.New(InitAws())

	input := &s3.ListObjectsV2Input{
    Bucket:  aws.String(GetAwsInfo("s3_bucket_dl")),
	}

	resp, err := client.ListObjectsV2(input)
  if err != nil {
    panic(err)
  }
  for _, c := range resp.Contents {
    fmt.Println(*c.Key)
  }
}
