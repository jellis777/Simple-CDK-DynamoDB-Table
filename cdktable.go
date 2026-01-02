package cdktable

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type CdktableStackProps struct {
	StackProps awscdk.StackProps
}

func NewCdktableStack(scope constructs.Construct, id string, props *CdktableStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	this := stack

	// The code that defines your stack goes here

	tableName := "barjokes-cdk"
	params := &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: aws.String("NAME"),
			Type: dynamodb.AttributeType_STRING,
		},
		BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   &tableName,
	}
	_ = dynamodb.NewTable(this, &tableName, params)

	return stack
}
