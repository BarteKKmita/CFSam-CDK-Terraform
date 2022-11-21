package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	provider.NewAwsProvider(stack, jsii.String("Provider"), &provider.AwsProviderConfig{})
	//var subnetInstance subnet.Subnet
	//var group securitygroup.SecurityGroupIngress
	//if "prod" == "prod" {
	//	newVpc := newVPC(stack)
	//	newSubnet(stack, newVpc)
	//	gateway := newInternetGateway(stack, newVpc)
	//	table := newRouteTable(stack, newVpc)
	//	newRoute(stack, table, gateway)
	//	newSecurityGroup()
	//}
	////newLambdaFunction(stack, subnetInstance, group)
	newDynamoDBTable(stack)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "terraform_cdk")

	app.Synth()
}
