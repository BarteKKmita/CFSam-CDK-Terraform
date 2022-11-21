package main

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/dynamodbtable"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func newDynamoDBTable(stack cdktf.TerraformStack) {
	dynamodbtable.NewDynamodbTable(stack, jsii.String("IaCPresentationDatabase"), &dynamodbtable.DynamodbTableConfig{
		BillingMode: jsii.String("PAY_PER_REQUEST"),
		Name:        jsii.String("IaCPresentationDatabase"),
		Attribute: map[string]interface{}{
			"id":       "S",
			"username": "S",
		},
		HashKey:  jsii.String("id"),
		RangeKey: jsii.String("username"),
	})
}
