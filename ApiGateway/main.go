// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var language = map[string]string{
	"en": "Hello",
	"cn": "NiHao",
}

type reponse struct {
	Message   string            `json:"message"`
	Info      map[string]string `json:"info"`
	TimeStamp time.Time         `json:"timestamp"`
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	info := request.QueryStringParameters
	lang := info["lang"]

	message := language[lang] + " " + name

	returnResponse := reponse{
		Message:   message,
		Info:      info,
		TimeStamp: time.Now(),
	}

	out, err := json.Marshal(returnResponse)
	fmt.Println(string(out))
	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(out),
	}, nil

}
