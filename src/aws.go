package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		Bucket: aws.String(GetAwsInfo("s3_bucket_dl")),
	}

	resp, err := client.ListObjectsV2(input)
	if err != nil {
		panic(err)
	}
	for _, c := range resp.Contents {
		fmt.Println(*c.Key)
	}
}

func Download(object_key string, filename string) error {
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(InitAws())

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", filename, err)
	}
	defer f.Close()

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(GetAwsInfo("s3_bucket_dl")),
		Key:    aws.String(object_key + "/" + filename),
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)

	return err
}
