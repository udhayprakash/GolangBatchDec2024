package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request is the structure for the input event
type Request struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

// Response is the structure for the output result
type Response struct {
	Sum float64 `json:"sum"`
}

// Handler is the Lambda function handler
func Handler(ctx context.Context, req Request) (Response, error) {
	sum := req.Num1 + req.Num2
	return Response{
		Sum: sum,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

/*

Installation go sdk

	go get github.com/aws/aws-lambda-go/lambda

Build the lambda function

	GOOS=linux GOARCH=amd64 go build -o main main.go


Local setup need environment variables

	Lambda environment variables [_LAMBDA_SERVER_PORT AWS_LAMBDA_RUNTIME_API]

create a deployment package
	chmod +x bootstrap
	zip deployment.zip bootstrap main


create the IAM Role

	aws iam create-role --role-name lambda-execution-role --assume-role-policy-document file://trust-policy.json

Attach the policy
	aws iam attach-role-policy --role-name lambda-execution-role --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole


Deploy to AWS Lambda:

	aws lambda create-function \
		--function-name golang-lambda-addition \
		--zip-file fileb://deployment.zip \
		--handler bootstrap \
		--runtime provided.al2 \
		--role arn:aws:iam::AWS_ACCOUNT_ID:role/myLambdaRole


aws lambda update-function-code \
    --function-name golang-lambda-addition \
    --zip-file fileb://deployment.zip
*/
