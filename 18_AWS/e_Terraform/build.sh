#!/bin/bash

# go mod init lambda-timezones
go get github.com/aws/aws-lambda-go/lambda


GOOS=linux GOARCH=amd64 go build -o main main.go

echo '#!/bin/sh' > bootstrap
echo './main' >> bootstrap
chmod +x bootstrap


zip deployment.zip bootstrap main

