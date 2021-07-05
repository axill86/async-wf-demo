# Resources repository :-)

This is a demo project which shows the workflow implementation based on AWS stack.
In order to install the actual dependencies run  
* `npm install`
# Structure
Available stacks can be viewed using following command
* `cdk list`  
That modules supposed to be installed in order they mentioned (at least for the first time)  
  
As for now that repository includes following modules:
  * *EventBusStack* - creates eventbus which can be used by producer/workflows 
  * *ProducerResourcesStack* - creates the producer. In that simple example that creates only IAM role which can put events in eventbridge.
  * *DummyWorkflowStack* - very simple workflow implemented via stepfunctions
  * *SmartWorkflowStack* - yet another workflow implementation
  * *ToolResourcesStack* - tool implementation which uses SQS to communicate
  * *CompleterStack* - cmd tool which listens for output queue of the tool and completes the step function   
  * *WorkflowRulesStack* - describes eventbus rules, which actually perform routing between destinations
  
The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

 * `npm run build`   compile typescript to js
 * `npm run watch`   watch for changes and compile
 * `npm run test`    perform the jest unit tests
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
