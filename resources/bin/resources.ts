#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import {EventBusStack} from '../lib/eventbus-stack';
import {DummyWorkflowStack} from "../lib/dummy-workflow-stack";
import {ProducerResourcesStack} from "../lib/producer-resources-stack";
import {WorkflowRulesStack} from "../lib/workflow-rules-stack";
import {SmartWorkflowStack} from "../lib/smart-workflow-stack";
import {ToolResourcesStack} from "../lib/tool-resources-stack";
import {CompleterStack} from "../lib/completer-stack";

const app = new cdk.App();
new EventBusStack(app, 'EventBusStack', {});
new DummyWorkflowStack(app, 'DummyWorkflowStack', {})
new ProducerResourcesStack(app, 'ProducerResourcesStack', {})
new WorkflowRulesStack(app, 'WorkflowRulesStack', {})
new SmartWorkflowStack(app, 'SmartWorkflowStack', {})
new ToolResourcesStack(app, 'ToolResourcesStack', {})
new CompleterStack(app, 'CompleterStack', {})