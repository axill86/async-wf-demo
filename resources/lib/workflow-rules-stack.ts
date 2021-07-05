import * as cdk from '@aws-cdk/core'
import * as ssm from '@aws-cdk/aws-ssm'
import * as events from '@aws-cdk/aws-events'
import * as targets from '@aws-cdk/aws-events-targets'
import * as sfn from '@aws-cdk/aws-stepfunctions'
import * as sqs from '@aws-cdk/aws-sqs'
import {AppConstants} from "./constants";
/**
 * This stack only contains rules for mapping events to workflows
 */
export class WorkflowRulesStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        const eventBusArnParam = ssm.StringParameter.fromStringParameterName(this, "eventbus-arn-param"
            , AppConstants.PARAM_EVENTBUS_ARN)
        const eventBus = events.EventBus.fromEventBusArn(this, "eventbus", eventBusArnParam.stringValue)
        const dummyWorkflowArnParam = ssm.StringParameter.fromStringParameterName(this, "dummy-workflow-arn-param"
            , AppConstants.PARAM_DUMMY_WORKFLOW_ARN)
        const dummyWorkflow = sfn.StateMachine.fromStateMachineArn(this, "dummy-workflow"
            , dummyWorkflowArnParam.stringValue)
        //just map the rule to the dummy
        const createDummyCalcRule = new events.Rule(this, "create-dummy-calc-rule", {
            eventBus: eventBus,
            targets: [new targets.SfnStateMachine(dummyWorkflow, {
                input: events.RuleTargetInput.fromEventPath("$.detail"),
            })],
            eventPattern: {
                detailType: ["CreateCalculation"],
                source: ["minicompute"],
                detail: {
                    flow: [{ "exists": false  }, "", "DUMMY"]
                }
            }
        })

        const smartWorkflowArn = ssm.StringParameter.fromStringParameterName(this, "smart-workflow-arn-param"
            , AppConstants.PARAM_SMART_WORKFLOW_ARN)
        const smartWorkflow = sfn.StateMachine.fromStateMachineArn(this, "smart-workflow"
            , smartWorkflowArn.stringValue)

        const createSmartCalcRule = new events.Rule(this, "create-smart-calc-rule", {
            eventBus: eventBus,
            targets: [new targets.SfnStateMachine(smartWorkflow, {
                input: events.RuleTargetInput.fromEventPath("$.detail")
            })],
            eventPattern: {
                detailType: ["CreateCalculation"],
                source: ["minicompute"],
                detail: {
                    flow: ["SMART"],
                }
            }
        })
        //rule for dummy workflow
        const toolQueueArnParam = ssm.StringParameter.fromStringParameterName(this, "tool-input-queue-arn-param", AppConstants.PARAM_TOOL_INPUT_QUEUE_ARN)
        const toolInputQueue = sqs.Queue.fromQueueArn(this, "tool-input-queue", toolQueueArnParam.stringValue)
        const toolInputTarget = new targets.SqsQueue(toolInputQueue, {
            message: events.RuleTargetInput.fromEventPath("$.detail")
        })
        const toolExecutionRule = new events.Rule(this, "execute-tool-rule", {
            eventBus: eventBus,
            targets: [toolInputTarget],
            eventPattern: {
                detailType: ["ToolRequest"],
                source: ["minicompute-smart-workflow"],
            }
        })
    }
}