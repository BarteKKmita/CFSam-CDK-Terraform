provider "aws" {
  region = "us-east-1"
}

locals {
  enviroment="prod"
}

resource "aws_iam_role" "iam_for_lambda" {
  name                = "iam_for_lambda"
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"]
  assume_role_policy  = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "LambdaPolicy" {
  name        = "LambdaPolicy"
  description = "A test lambda policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Action": [
        "events:PutEvents"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "LambdaPolicyAttachment" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.LambdaPolicy.arn
}

resource "aws_lambda_function" "test_prod_lambda" {
  filename         = "main.zip"
  function_name    = "lambda_function_name"
  handler          = "main"
  source_code_hash = filebase64sha256("main.zip")
  memory_size      = 128
  timeout          = 10
  runtime          = "go1.x"
  role             = aws_iam_role.iam_for_lambda.arn
  vpc_config {
    security_group_ids = [aws_security_group.IaCPresentationFunctionSecurityGroup.id]
    subnet_ids         = [aws_subnet.IacPresentationPublicSubnetTwo.id]
  }
}

