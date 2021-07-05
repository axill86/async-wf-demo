package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSReporterServiceConfig struct {
	QueueUrl string
	Aws      aws.Config
}

type SQSReporterRequest struct {
	ToolResponse
}

type SQSReporterResponse struct {
	MessageId string
}
type SQSReporterService interface {
	SendEvent(ctx context.Context, request SQSReporterRequest) (*SQSReporterResponse, error)
}

type sqsReporterServiceImpl struct {
	sqs      *sqs.Client
	QueueUrl string
}

func NewSQSReporterService(config SQSReporterServiceConfig) SQSReporterService {
	return &sqsReporterServiceImpl{
		sqs:      sqs.NewFromConfig(config.Aws),
		QueueUrl: config.QueueUrl,
	}
}
func (s *sqsReporterServiceImpl) SendEvent(ctx context.Context, request SQSReporterRequest) (*SQSReporterResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request %e", err)
	}
	res, err := s.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(s.QueueUrl),
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request to queue %e", err)
	}
	return &SQSReporterResponse{MessageId: *res.MessageId}, nil
}
