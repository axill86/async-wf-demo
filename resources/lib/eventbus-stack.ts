import * as cdk from '@aws-cdk/core';
import * as events from '@aws-cdk/aws-events'
import * as ssm from '@aws-cdk/aws-ssm'
import {AppConstants} from "./constants";

/*
This stack only creates eventbus.
 */
export class EventBusStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
    const eventBus = new events.EventBus(this, "minicompute-events")
    const eventBusNameParam = new ssm.StringParameter(this, "minicompute-evnets-name", {
      parameterName: AppConstants.PARAM_EVENTBUS_NAME,
      stringValue: eventBus.eventBusName
    })
    const eventBusArnParam = new ssm.StringParameter(this, "eventbus-arn-param", {
      parameterName: AppConstants.PARAM_EVENTBUS_ARN,
      stringValue: eventBus.eventBusArn
    })
  }
}
