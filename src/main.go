package main

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrUndefinedHTTPMethod is thrown when a name is not provided
	ErrUndefinedHTTPMethod = errors.New("method not allowed")
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	switch request.HTTPMethod {
	case "GET":
		return events.APIGatewayProxyResponse{
			Body:            "Hello test 3, method: GET, TEST: " + os.Getenv("TEST") + ", GLOBAL: " + os.Getenv("GLOBAL"),
			StatusCode:      200,
			IsBase64Encoded: false,
		}, nil

	case "POST":
		return events.APIGatewayProxyResponse{
			Body:            "Hello test, method POST, " + os.Getenv("TEST") + ", GLOBAL: " + os.Getenv("GLOBAL") + "\nBody:\n\n" + request.Body,
			StatusCode:      200,
			IsBase64Encoded: false,
		}, nil

	}

	return events.APIGatewayProxyResponse{}, ErrUndefinedHTTPMethod

}

func main() {
	lambda.Start(Handler)
}
