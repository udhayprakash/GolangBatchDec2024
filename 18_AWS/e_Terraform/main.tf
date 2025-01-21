provider "aws" {
  region = "us-east-1"  # Replace with your preferred AWS region
}

resource "aws_lambda_function" "hello_world" {
  function_name = "hello-world-lambda"
  handler       = "bootstrap"
  runtime       = "provided.al2"
  role          = aws_iam_role.lambda_exec2.arn
  filename      = "deployment.zip"

  source_code_hash = filebase64sha256("deployment.zip")

  environment {
    variables = {
      FOO = "bar"
    }
  }
}

resource "aws_iam_role" "lambda_exec2" {
  name = "lambda_exec2_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_exec2_policy" {
  role       = aws_iam_role.lambda_exec2.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

output "lambda_function_name" {
  value = aws_lambda_function.hello_world.function_name
}