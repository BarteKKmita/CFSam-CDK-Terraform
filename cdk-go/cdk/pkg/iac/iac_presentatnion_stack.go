package iac

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/constructs-go/constructs/v10"
)

type CdkGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoIaCPresentationStack(scope constructs.Construct, id string, props *CdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	var vpc awsec2.Vpc
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	if "prod" == "prod" {
		vpc = newVPC(stack)
	}

	function := newLambda(stack, vpc)
	newLambdaPermission(stack, function)
	api := newApiV2(stack)
	integration := newApiIntegration(stack, api, function)
	newApiRoute(stack, api, integration)
	newApiStage(stack, api)
	awseventstargets.NewLambdaFunction(function, nil)

	newDynamoTable(stack)

	return stack
}
