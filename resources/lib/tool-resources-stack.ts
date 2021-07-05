import * as cdk from '@aws-cdk/core'
import * as sqs from '@aws-cdk/aws-sqs'
import * as iam from '@aws-cdk/aws-iam'
import * as ssm from '@aws-cdk/aws-ssm'
import {AppConstants} from "./constants";
import {PolicyStatement} from "@aws-cdk/aws-iam";

//This stack will create needed for the tool resources (sqs queues and iam role)
export class ToolResourcesStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        const inputQueue = new sqs.Queue(this, "input-queue", )
        const inputQueuePolicy = new sqs.QueuePolicy(this, "input-queue-policy", {
            queues: [inputQueue],
            }
        )
        inputQueuePolicy.document.addStatements(
            new PolicyStatement({
                principals: [new iam.ServicePrincipal('events.amazonaws.com')],
                actions: ['sqs:SendMessage'],
                resources: [inputQueue.queueArn]
            })
        )
        const outputQueue = new sqs.Queue(this, "output-queue")
        const toolRole = new iam.Role(this, "tool-role", {
            assumedBy: new iam.AccountRootPrincipal()
        })
        inputQueue.grantConsumeMessages(toolRole)
        outputQueue.grantSendMessages(toolRole)
        const inputQueueUrlOutputParam = new cdk.CfnOutput(this, "input-queue-url", {
            exportName: "input-queue-url",
            value: inputQueue.queueUrl
        })

        const outputQueueUrlParam = new cdk.CfnOutput(this, "output-queue-url", {
            exportName: "output-queue-url",
            value: outputQueue.queueUrl
        })

        const toolRoleParam = new cdk.CfnOutput(this, "tool-role-name", {
            exportName: "tool-role",
            value: toolRole.roleArn
        })

        const inputQueueArnParam = new ssm.StringParameter(this, "input-queue-arn-param", {
            parameterName: AppConstants.PARAM_TOOL_INPUT_QUEUE_ARN,
            stringValue: inputQueue.queueArn
        })

        const outputQueueArnParam = new ssm.StringParameter(this, "output-queue-arn-param", {
            parameterName: AppConstants.PARAM_TOOL_OUTPUT_QUEUE_ARN,
            stringValue: outputQueue.queueArn
        })
    }
}