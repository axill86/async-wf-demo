import * as cdk from '@aws-cdk/core'
import * as ssm from '@aws-cdk/aws-ssm'
import * as iam from '@aws-cdk/aws-iam'
import * as sqs from '@aws-cdk/aws-sqs'
import * as sfn from '@aws-cdk/aws-stepfunctions'
import {AppConstants} from "./constants";

export class CompleterStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        const queueArnParam = ssm.StringParameter.fromStringParameterName(this, "queue-arn-param", AppConstants.PARAM_TOOL_OUTPUT_QUEUE_ARN)
        const queue = sqs.Queue.fromQueueArn(this, "queue", queueArnParam.stringValue)

        const workflowArnParam = ssm.StringParameter.fromStringParameterName(this, "workflow-arn-param", AppConstants.PARAM_SMART_WORKFLOW_ARN)
        const workflow = sfn.StateMachine.fromStateMachineArn(this, "workfow", workflowArnParam.stringValue)
        const completerRole = new iam.Role(this, "completer-role", {
            assumedBy: new iam.AccountRootPrincipal()
        })
        queue.grantConsumeMessages(completerRole)
        workflow.grantTaskResponse(completerRole)
        const roleNameOutput = new cdk.CfnOutput(this, "role-arn-output", {
            exportName: "completer-role",
            value: completerRole.roleArn
        })
    }
}