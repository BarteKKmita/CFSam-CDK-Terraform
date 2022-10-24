package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(ctx context.Context) error {
	log.Println("Hello from CDK!!!")
	return nil
}

func main() {
	lambda.Start(handler)
}
