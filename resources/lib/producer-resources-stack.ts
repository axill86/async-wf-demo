import * as cdk from '@aws-cdk/core'
import * as iam from '@aws-cdk/aws-iam'
import * as ssm from '@aws-cdk/aws-ssm'
import * as events from '@aws-cdk/aws-events'
import {AppConstants} from "./constants";

/**
 * Creates iam role for producer with only permission to write events
 */
export class ProducerResourcesStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);
        const eventBusArnParam = ssm.StringParameter.fromStringParameterName(this, "eventbus-arn-param",
            AppConstants.PARAM_EVENTBUS_ARN)
        const eventBus = events.EventBus.fromEventBusArn(this, "evenbus", eventBusArnParam.stringValue)
        const producerRole = new iam.Role(this, "producer-role", {
            assumedBy: new iam.AccountRootPrincipal()
        })
        eventBus.grantPutEventsTo(producerRole)
        const roleOutput = new cdk.CfnOutput(this, "producer-role-output", {
            exportName: "producer-role",
            value: producerRole.roleArn
        })
    }
}