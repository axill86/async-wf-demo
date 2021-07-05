package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type CompleteTaskRequest struct {
	StatusMessage string `json:"status_message"`
	RequestId     string `json:"request_id"`
}

type CompleterConfig struct {
	AWS          aws.Config
}

type Completer interface {
	CompleteTask(ctx context.Context, request CompleteTaskRequest) error
}

type completerImpl struct {
	client *sfn.Client
}

func (c *completerImpl) CompleteTask(ctx context.Context, request CompleteTaskRequest)  error {
	body, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("error serializing json %e", err)
	}
	_, err = c.client.SendTaskSuccess(ctx, &sfn.SendTaskSuccessInput{
		Output:    aws.String(string(body)),
		TaskToken: aws.String(request.RequestId),
	})
	if err != nil {
		return fmt.Errorf("error sending callback %e", err)
	}
	return nil
}

func NewCompleter(config CompleterConfig) Completer {
	return &completerImpl{
		client: sfn.NewFromConfig(config.AWS),
	}
}
