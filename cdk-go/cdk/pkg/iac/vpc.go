package iac

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func newVPC(stack awscdk.Stack) awsec2.Vpc {
	return awsec2.NewVpc(stack, jsii.String("IaCPresentationVpc"), &awsec2.VpcProps{
		Cidr:               jsii.String("10.0.0.0/16"),
		EnableDnsHostnames: jsii.Bool(true),
		EnableDnsSupport:   jsii.Bool(true),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			newPrivateSubnetGroup(),
			newPublicSubnetGroup(),
		},
		MaxAzs:      jsii.Number(2),
		NatGateways: jsii.Number(1),
		VpcName:     jsii.String("IaCPresentationVpc"),
	})
}

func newPrivateSubnetGroup() *awsec2.SubnetConfiguration {
	return &awsec2.SubnetConfiguration{
		Name:       jsii.String("IacPresentationPrivate"),
		SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
		CidrMask:   jsii.Number(24),
		Reserved:   jsii.Bool(false),
	}
}

func newPublicSubnetGroup() *awsec2.SubnetConfiguration {
	return &awsec2.SubnetConfiguration{
		Name:                jsii.String("IacPresentationPublic"),
		SubnetType:          awsec2.SubnetType_PUBLIC,
		CidrMask:            jsii.Number(24),
		MapPublicIpOnLaunch: jsii.Bool(false),
		Reserved:            jsii.Bool(false),
	}
}
