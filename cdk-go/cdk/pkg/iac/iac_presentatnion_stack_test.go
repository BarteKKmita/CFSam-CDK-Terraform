package iac

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkGoStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)
	// WHEN
	stack := NewCdkGoIaCPresentationStack(app, "CdkGoStack", nil)
	template := assertions.Template_FromStack(stack, nil)
	// THEN
	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 300,
	})
	template.HasResourceProperties(jsii.String("AWS::Events::Rule"), map[string]interface{}{
		"ScheduleExpression": "cron(0/15 * ? * * *)",
	})

}
