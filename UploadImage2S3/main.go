// main.go
package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"net/http"
)

var s3session *s3.S3

const (
	REGION     = "us-west-1"
	BucketName = "image-lambda-xixian"
)

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}

func Handler(event InputEvent) (string, error) {
	image := GetImage(event.Link)
	_, err := s3session.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(image),
		Bucket: aws.String(BucketName),
		Key:    aws.String(event.Key),
	})
	if err != nil {
		return "Error", err
	}
	return "ok", err
}

// GetImage  Get image by url and transform to byte slice
func GetImage(url string) (bytes []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}
