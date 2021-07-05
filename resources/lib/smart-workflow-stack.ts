import * as cdk from '@aws-cdk/core'
import * as sfn from '@aws-cdk/aws-stepfunctions'
import * as ssm from '@aws-cdk/aws-ssm'
import * as events from '@aws-cdk/aws-events'
import {AppConstants} from "./constants"

/**
 * This workflow will execute the tool
 */
export class SmartWorkflowStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

        const eventBusArnParam = ssm.StringParameter.fromStringParameterName(this, "eventbus-arn-param", AppConstants.PARAM_EVENTBUS_ARN)
        const eventBus = events.EventBus.fromEventBusArn(this, "event-bus", eventBusArnParam.stringValue)
        //tool integration
        const toolStep = new sfn.CustomState(this, "execute-tool-state", {
            stateJson: {
                Type: "Task",
                Resource: 'arn:aws:states:::events:putEvents.waitForTaskToken',
                TimeoutSeconds: 300,
                Parameters: {
                    Entries: [
                        {
                            Detail: {
                                'request_id.$': '$$.Task.Token',
                                'initiator.$': '$$.Execution.Input.initiator',
                                'order.$': '$$.Execution.Input.order'
                            },
                            DetailType: 'ToolRequest',
                            Source: 'minicompute-smart-workflow',
                            EventBusName: eventBus.eventBusName
                        }
                    ]
                }
            }
        })

        const workflow = new sfn.StateMachine(this, "smart-workflow", {
            definition: toolStep
        })

        //grant put events to workflow
        eventBus.grantPutEventsTo(workflow)

        const workflowNameParam = new ssm.StringParameter(this, "workflow-name-param", {
            parameterName: AppConstants.PARAM_SMART_WORKFLOW_NAME,
            stringValue: workflow.stateMachineName
        })
        const workflowArnParam = new ssm.StringParameter(this, "workflow-arn-param", {
            parameterName: AppConstants.PARAM_SMART_WORKFLOW_ARN,
            stringValue: workflow.stateMachineArn
        })
    }
}