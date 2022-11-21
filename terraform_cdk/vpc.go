package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/internetgateway"
	"cdk.tf/go/stack/generated/hashicorp/aws/route"
	"cdk.tf/go/stack/generated/hashicorp/aws/routetable"
	"cdk.tf/go/stack/generated/hashicorp/aws/securitygroup"
	"cdk.tf/go/stack/generated/hashicorp/aws/subnet"
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func newRoute(stack cdktf.TerraformStack, table routetable.RouteTable, gateway internetgateway.InternetGateway) route.Route {
	return route.NewRoute(stack, jsii.String("IacPresentationRoute"), &route.RouteConfig{
		RouteTableId:         table.Id(),
		GatewayId:            gateway.Id(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
	})
}

func newRouteTable(stack cdktf.TerraformStack, newVpc vpc.Vpc) routetable.RouteTable {
	return routetable.NewRouteTable(stack, jsii.String("IacPresentationRouteTable"), &routetable.RouteTableConfig{
		VpcId: newVpc.Id(),
	})
}

func newInternetGateway(stack cdktf.TerraformStack, newVpc vpc.Vpc) internetgateway.InternetGateway {
	return internetgateway.NewInternetGateway(stack, jsii.String("IacPresentationInternetGateway"), &internetgateway.InternetGatewayConfig{
		VpcId: newVpc.Id(),
	})
}

func newSubnet(stack cdktf.TerraformStack, newVpc vpc.Vpc) subnet.Subnet {
	return subnet.NewSubnet(stack, jsii.String("IacPresentationPublicSubnetOne"), &subnet.SubnetConfig{
		CidrBlock:        jsii.String("10.0.0.0/24"),
		AvailabilityZone: jsii.String("us-east-1a"),
		VpcId:            newVpc.Id(),
	})
}

func newSecurityGroup() securitygroup.SecurityGroupIngress {
	return securitygroup.SecurityGroupIngress{
		CidrBlocks: &[]*string{
			jsii.String("0.0.0.0/0"),
		},
		FromPort: jsii.Number(0),
		ToPort:   jsii.Number(65535),
		Protocol: jsii.String("tcp"),
	}
}

func newVPC(stack cdktf.TerraformStack) vpc.Vpc {

	return vpc.NewVpc(stack, jsii.String("IaCPresentationVpc"), &vpc.VpcConfig{
		CidrBlock:          jsii.String("10.0.0.0/16"),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
		InstanceTenancy:    jsii.String("default"),
	})
}
