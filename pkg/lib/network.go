package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	awsalb "github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type Network struct {
	Vpc awsec2.IVpc
	Alb awsalb.IApplicationLoadBalancer
}

func NewNetwork(scope constructs.Construct, id *string) *Network {

	this := constructs.NewConstruct(scope, id, nil)

	vpc := awsec2.NewVpc(this, jsii.String("vpc"), nil)
	alb := awsalb.NewApplicationLoadBalancer(this, jsii.String("alb"), &awsalb.ApplicationLoadBalancerProps{
		Vpc:            vpc,
		InternetFacing: jsii.Bool(true),
	})

	return &Network{
		Vpc: vpc,
		Alb: alb,
	}
}
