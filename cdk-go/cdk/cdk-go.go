package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoStack(scope constructs.Construct, id string, props *CdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	function := awslambda.NewFunction(stack, jsii.String("myGoLambda"), &awslambda.FunctionProps{
		Description:  jsii.String("Hello function"),
		FunctionName: jsii.String("Hello"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.Code_FromAsset(jsii.String("../cmd/dist/main.zip"), nil),
	})

	lambdaFunction := awseventstargets.NewLambdaFunction(function, nil)
	var targets []awsevents.IRuleTarget
	targets = append(targets, lambdaFunction)
	awsevents.NewRule(stack, jsii.String("schelduedRule"), &awsevents.RuleProps{
		Description: jsii.String("Schedule rule"),
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Minute:  jsii.String("0/15"),
			Hour:    jsii.String("*"),
			WeekDay: jsii.String("*"),
			Month:   jsii.String("*"),
			Year:    jsii.String("*"),
		}),
		Targets: &targets,
	})

	awssqs.NewQueue(stack, jsii.String("CdkGoQueue"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	return stack
}

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)

	NewCdkGoStack(app, "CdkGoStack", &CdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
