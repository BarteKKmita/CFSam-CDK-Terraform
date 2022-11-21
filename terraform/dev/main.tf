
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

resource "aws_lambda_function" "test_dev_lambda" {
  filename         = "main.zip"
  function_name    = "lambda_function_name"
  handler          = "main"
  source_code_hash = filebase64sha256("main.zip")
  memory_size      = 128
  timeout          = 10
  runtime          = "go1.x"
  role             = aws_iam_role.iam_for_lambda.arn
}

resource "aws_apigatewayv2_api" "IaCPresentationApi" {
  name          = "IaCPresentationApi"
  protocol_type = "HTTP"
  cors_configuration {
    allow_methods = ["*"]
    allow_origins = ["*"]
  }
}

resource "aws_apigatewayv2_integration" "IaCPresentationApiIntegration" {
  api_id        = aws_apigatewayv2_api.IaCPresentationApi.id
  integration_uri = aws_lambda_function.test_prod_lambda[0].arn
  payload_format_version = "2.0"
  integration_method = "POST"
  integration_type = "AWS_PROXY"
}

resource "aws_apigatewayv2_stage" "IaCPresentationApiStage" {
  api_id = aws_apigatewayv2_api.IaCPresentationApi.id
  name   = "testStage"
  auto_deploy = true
}

resource "aws_apigatewayv2_route" "IaCPresentationApiLambdaRoute" {
  api_id         = aws_apigatewayv2_api.IaCPresentationApi.id
  route_key      = "GET /hello"
  target = "integrations/${aws_apigatewayv2_integration.IaCPresentationApiIntegration.id}"
}

resource "aws_lambda_permission" "IaCPresentationApiLambdaPermission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_prod_lambda[0].function_name
  principal     = "apigateway.amazonaws.com"
}


