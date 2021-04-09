package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type ECS struct {
	Cluster awsecs.ICluster
	TaskDef awsecs.ITaskDefinition
	Service awsecs.IService
}

type ECSProps struct {
	Vpc awsec2.IVpc
}

func NewECS(scope constructs.Construct, id *string, props *ECSProps) *ECS {

	this := constructs.NewConstruct(scope, id, nil)

	cluster := awsecs.NewCluster(this, jsii.String("cluster"), &awsecs.ClusterProps{
		ContainerInsights: jsii.Bool(true),
		Vpc:               props.Vpc,
	})

	taskDef := awsecs.NewTaskDefinition(this, jsii.String("taskDef"), nil)
	taskDef.AddContainer(jsii.String("nginx"), &awsecs.ContainerDefinitionOptions{
		Image: awsecs.ContainerImage_FromRegistry(jsii.String("nginx/nginx:latest"), nil),
	})

	service := awsecs.NewFargateService(this, jsii.String("service"), &awsecs.FargateServiceProps{
		Cluster:         cluster,
		TaskDefinition:  taskDef,
		PlatformVersion: awsecs.FargatePlatformVersion_VERSION1_4,
	})

	return &ECS{
		Cluster: cluster,
		TaskDef: taskDef,
		Service: service,
	}
}
