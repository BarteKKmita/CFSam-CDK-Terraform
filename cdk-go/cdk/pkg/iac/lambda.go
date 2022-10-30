package iac

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func newLambda(stack awscdk.Stack, vpc awsec2.Vpc) awslambda.Function {
	return awslambda.NewFunction(stack, jsii.String("myGoLambda"), &awslambda.FunctionProps{
		Description:  jsii.String("Hello function"),
		FunctionName: jsii.String("Hello"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.Code_FromAsset(jsii.String("../lambda/cmd/dist/main.zip"), nil),
		Vpc:          vpc,
	})
}

func newLambdaPermission(stack awscdk.Stack, function awslambda.Function) {
	awslambda.NewCfnPermission(stack, jsii.String("ApiGatewayPermission"), &awslambda.CfnPermissionProps{
		Action:       jsii.String("lambda:InvokeFunction"),
		Principal:    jsii.String("apigateway.amazonaws.com"),
		FunctionName: function.FunctionName(),
	})
}
