package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"tool/config"
	"tool/service"
)

//LambdaSQSHandler handles input events via just sending response back to output queue
func LambdaSQSHandler(ctxt context.Context, event events.SQSEvent) error {
	conf := config.ReadSQSConfig()
	toolService := service.NewToolService(service.ToolServiceConfig{})
	awsClient := aws.NewConfig()
	reporterService := service.NewSQSReporterService(service.SQSReporterServiceConfig{
		QueueUrl: conf.OutputQueue,
		Aws:      *awsClient,
	})
	for _, record := range event.Records {
		body := record.Body
		var request service.ToolRequest
		if err := json.Unmarshal([]byte(body), &request); err != nil {
			return fmt.Errorf("error parsing message body %e", err)
		}
		result, err := toolService.ServeRequest(ctxt, request)
		if err != nil {
			return fmt.Errorf("service error %e", err)
		}
		log.Printf("Tool executed with input %v and result %v", request, result)
		resp, err := reporterService.SendEvent(ctxt, service.SQSReporterRequest{ToolResponse: *result})
		if err != nil {
			return fmt.Errorf("error sending notification %e", err)
		}
		log.Printf("Sent notification with result %v", resp)
	}
	return nil
}
