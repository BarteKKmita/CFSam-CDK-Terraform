package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/subnet"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/lambdafunction"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/securitygroup"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func newLambdaFunctionVPC(stack cdktf.TerraformStack, vpcSubnet subnet.Subnet, group securitygroup.SecurityGroup) {
	lambdafunction.NewLambdaFunction(stack, jsii.String("IaCPresentationFunction"), &lambdafunction.LambdaFunctionConfig{
		Description:  jsii.String("Hello Function"),
		FunctionName: jsii.String("Hello"),
		MemorySize:   jsii.Number(128),
		Timeout:      jsii.Number(10),
		Runtime:      jsii.String("go1.x"),
		Handler:      jsii.String("main"),
		Filename:     jsii.String("../lambda/cmd/dist/main.zip"), // TODO
		VpcConfig: &lambdafunction.LambdaFunctionVpcConfig{
			SecurityGroupIds: &[]*string{
				group.Id(),
			},
			SubnetIds: &[]*string{
				vpcSubnet.Id(),
			},
		},
	})
}

//func newLambdaFunction(stack cdktf.TerraformStack, vpcSubnet subnet.Subnet, group securitygroup.SecurityGroupIngress) {
//	var lambdaVpcConfig lambdafunction.LambdaFunctionVpcConfig
//	if "prod" == "prod" {
//		lambdaVpcConfig = lambdafunction.LambdaFunctionVpcConfig{
//			SecurityGroupIds: &[]*string{
//				group.Id(),
//			},
//			SubnetIds: &[]*string{
//				vpcSubnet.Id(),
//			},
//		}
//	}
//	lambdafunction.NewLambdaFunction(stack, jsii.String("IaCPresentationFunction"), &lambdafunction.LambdaFunctionConfig{
//		Description:  jsii.String("Hello Function"),
//		FunctionName: jsii.String("Hello"),
//		MemorySize:   jsii.Number(128),
//		Timeout:      jsii.Number(10),
//		Runtime:      jsii.String("go1.x"),
//		Handler:      jsii.String("main"),
//		Filename:     jsii.String("../lambda/cmd/dist/main.zip"), // TODO
//		VpcConfig:    &lambdaVpcConfig,
//	})
//}
