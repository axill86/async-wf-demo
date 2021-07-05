# Here are resources for demo project of step functions with async workflow

## Structure

* [resources](./resources) - AWS resources needed for demo, created by AWS CDK
* [producer](./producer) - Command-line tool which generates events
* [tool](./tool) - Tool which listens to sqs queue and pushes the result out
* [task-completer](./task-completer) - That tool listens for messages from output queue of the tool and completes the step-function
