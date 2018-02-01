package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// MyEvent イベント
type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

// MyResponse レスポンス
type MyResponse struct {
	Message string `json:"Answer:"`
}

// HandleLambdaEvent ハンドラー
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	log.Print("logging test")

	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("examplebucket"),
	}
	result, _ := svc.ListObjectsV2(input)
	myObjects := result.Contents
	log.Print(myObjects)
	return MyResponse{Message: fmt.Sprintf("%s is %d yoears old!!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
