package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	awsalb "github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoEcsStackProps struct {
	awscdk.StackProps
}

func NewCdkGoEcsStack(scope constructs.Construct, id string, props *CdkGoEcsStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := awsec2.NewVpc(stack, jsii.String("vpc"), nil)

	cluster := awsecs.NewCluster(stack, jsii.String("cluster"), &awsecs.ClusterProps{
		ContainerInsights: jsii.Bool(true),
		Vpc:               vpc,
	})
	taskDef := awsecs.NewFargateTaskDefinition(stack, jsii.String("taskDef"), &awsecs.FargateTaskDefinitionProps{
		Cpu: jsii.Number(1024),
	})
	taskDef.AddContainer(jsii.String("nginx"), &awsecs.ContainerDefinitionOptions{
		Image: awsecs.ContainerImage_FromRegistry(jsii.String("nginx/nginx:latest"), nil),
	})
	service := awsecs.NewFargateService(stack, jsii.String("service"), &awsecs.FargateServiceProps{
		Cluster:         cluster,
		TaskDefinition:  taskDef,
		PlatformVersion: awsecs.FargatePlatformVersion_VERSION1_4,
		DesiredCount:    jsii.Number(1),
	})

	alb := awsalb.NewApplicationLoadBalancer(stack, jsii.String("alb"), &awsalb.ApplicationLoadBalancerProps{
		Vpc:            vpc,
		InternetFacing: jsii.Bool(true),
	})
	listener := alb.AddListener(jsii.String("listener"), &awsalb.BaseApplicationListenerProps{
		Port: jsii.Number(80),
	})
	listener.AddTargets(jsii.String("targets"), &awsalb.AddApplicationTargetsProps{
		Port:    jsii.Number(80),
		Targets: &[]awsalb.IApplicationLoadBalancerTarget{service},
	})

	//network := lib.NewNetwork(scope, jsii.String("network"))
	//lib.NewECS(scope, jsii.String("ecs"), &lib.ECSProps{Vpc: network.Vpc})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkGoEcsStack(app, "CdkGoEcsStack", &CdkGoEcsStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
