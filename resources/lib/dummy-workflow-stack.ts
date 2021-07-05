import * as cdk from '@aws-cdk/core'
import * as ssm from '@aws-cdk/aws-ssm'
import * as sfn from '@aws-cdk/aws-stepfunctions'
import {AppConstants} from "./constants";
/**
 * That one is with very simple workflow (dummy one)
 */
export class DummyWorkflowStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        const dummyStep = new sfn.Pass(this, "dummy-step")
        const dummyWorkflow = new sfn.StateMachine(this, "dummy-workflow", {
            definition: dummyStep
        })
        const dummyWorkflowArnParam = new ssm.StringParameter(this, "workflow-arn-param", {
            parameterName: AppConstants.PARAM_DUMMY_WORKFLOW_ARN,
            stringValue: dummyWorkflow.stateMachineArn
        })
        const dummyWorkflowNameParam = new ssm.StringParameter(this, "workflow-name-param", {
            parameterName: AppConstants.PARAM_DUMMY_WORKFLOW_NAME,
            stringValue: dummyWorkflow.stateMachineName
        })
    }
}