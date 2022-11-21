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
integration_uri = aws_lambda_function.test_prod_lambda.arn
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
function_name = aws_lambda_function.test_prod_lambda.function_name
principal     = "apigateway.amazonaws.com"
}



