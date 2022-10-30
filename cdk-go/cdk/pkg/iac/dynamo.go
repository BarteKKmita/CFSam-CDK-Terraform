package iac

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

func newDynamoTable(stack awscdk.Stack) {
	awsdynamodb.NewCfnTable(stack, jsii.String("IaCPresentationDatabase"), &awsdynamodb.CfnTableProps{
		KeySchema: []interface{}{&awsdynamodb.CfnTable_KeySchemaProperty{
			AttributeName: jsii.String("id"),
			KeyType:       jsii.String("HASH"),
		},
			&awsdynamodb.CfnTable_KeySchemaProperty{
				AttributeName: jsii.String("username"),
				KeyType:       jsii.String("RANGE"),
			},
		},
		AttributeDefinitions: []interface{}{
			awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("username"),
				AttributeType: jsii.String("S"),
			},
			awsdynamodb.CfnTable_AttributeDefinitionProperty{
				AttributeName: jsii.String("id"),
				AttributeType: jsii.String("S"),
			},
		},
		BillingMode: jsii.String("PAY_PER_REQUEST"),
		TableName:   jsii.String("IaCPresentationTable"),
	})

	//awsdynamodb.NewTable(stack, jsii.String("IaCPresentationDatabase"), &awsdynamodb.TableProps{
	//	PartitionKey: &awsdynamodb.Attribute{
	//		Name: jsii.String("id"),
	//		Type: awsdynamodb.AttributeType_STRING,
	//	},
	//	SortKey: &awsdynamodb.Attribute{
	//		Name: jsii.String("username"),
	//		Type: awsdynamodb.AttributeType_STRING,
	//	},
	//	//KeySchema:            []interface{}{string(keySchemaBytes)},
	//	//AttributeDefinitions: []interface{}{string(attributeDefinitionsBytes)},
	//	BillingMode: "PAY_PER_REQUEST",
	//	TableName:   jsii.String("IaCPresentationTable2"),
	//})
}
