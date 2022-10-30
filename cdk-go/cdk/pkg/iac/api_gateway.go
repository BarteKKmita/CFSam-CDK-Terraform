package iac

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func newApiV2(stack awscdk.Stack) awsapigatewayv2.CfnApi {
	return awsapigatewayv2.NewCfnApi(stack, jsii.String("IaCPresentationApi"), &awsapigatewayv2.CfnApiProps{
		ProtocolType: jsii.String("HTTP"),
		Name:         jsii.String("HelloApi"),
	})
}

func newApiIntegration(stack awscdk.Stack, api awsapigatewayv2.CfnApi, function awslambda.Function) awsapigatewayv2.CfnIntegration {
	integration := awsapigatewayv2.NewCfnIntegration(stack, jsii.String("Lambda Integration"), &awsapigatewayv2.CfnIntegrationProps{
		ApiId:                api.Ref(),
		IntegrationType:      jsii.String("AWS_PROXY"),
		IntegrationMethod:    jsii.String("Post"),
		IntegrationUri:       function.FunctionArn(),
		PayloadFormatVersion: jsii.String("1.0"),
	})
	return integration
}

func newApiRoute(stack awscdk.Stack, api awsapigatewayv2.CfnApi, integration awsapigatewayv2.CfnIntegration) awsapigatewayv2.CfnRoute {
	return awsapigatewayv2.NewCfnRoute(stack, jsii.String("Lambda Route"), &awsapigatewayv2.CfnRouteProps{
		ApiId:    api.Ref(),
		RouteKey: jsii.String("GET /hello"),
		Target:   jsii.String("integrations/" + *integration.Ref()),
	})
}

func newApiStage(stack awscdk.Stack, api awsapigatewayv2.CfnApi) awsapigatewayv2.CfnStage {
	return awsapigatewayv2.NewCfnStage(stack, jsii.String("testStage"), &awsapigatewayv2.CfnStageProps{
		ApiId:      api.Ref(),
		StageName:  jsii.String("testStage"),
		AutoDeploy: true,
	})
}
