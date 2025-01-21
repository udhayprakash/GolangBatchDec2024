#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 {create|update}"
    exit 1
}

# Check if an argument is provided
if [ $# -ne 1 ]; then
    usage
fi

# go mod init lambda-timezones
go get github.com/aws/aws-lambda-go/lambda


GOOS=linux GOARCH=amd64 go build -o main main.go

echo '#!/bin/sh' > bootstrap
echo './main' >> bootstrap
chmod +x bootstrap


zip deployment.zip bootstrap main

# Deploy to AWS Lambda
if [ "$1" == "create" ]; then
	aws lambda create-function \
		--function-name lambda-timezones \
		--zip-file fileb://deployment.zip \
		--handler bootstrap \
		--runtime provided.al2 \
		--role arn:aws:iam::AWS_ACCOUNT_ID:role/myLambdaRole 
elif [ "$1" == "update" ]; then
    aws lambda update-function-code \
        --function-name lambda-timezones \
        --zip-file fileb://deployment.zip
fi
