package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello")

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
