package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/google/uuid"
)

const (
	FlowDummy = "DUMMY"
	FlowSmart = "SMART"
)

type CreateCalculationRequest struct {
	Initiator string `json:"initiator"`
	Order     int    `json:"order"`
	Flow      string `json:"flow,omitempty"`
}

type CreateCalculationResponse struct {
	Id string
}

type CalculationServiceConfig struct {
	Aws          aws.Config
	EventBusName string
}
type CalculationService interface {
	CreateCalculation(ctx context.Context, request CreateCalculationRequest) (*CreateCalculationResponse, error)
}

type calculationServiceImpl struct {
	client       *eventbridge.Client
	eventBusName string
}

type eventData struct {
	Id string `json:"id"`
	CreateCalculationRequest
}

func (c *calculationServiceImpl) CreateCalculation(ctx context.Context, request CreateCalculationRequest) (*CreateCalculationResponse, error) {
	event := eventData{
		Id:                       uuid.New().String(),
		CreateCalculationRequest: request,
	}
	requestBody, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("error marshalling event data %e", err)
	}
	_, err = c.client.PutEvents(ctx, &eventbridge.PutEventsInput{
		Entries: []types.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(requestBody)),
				DetailType:   aws.String("CreateCalculation"),
				EventBusName: aws.String(c.eventBusName),
				Source:       aws.String("minicompute"),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request to eventbridge %e", err)
	}
	return &CreateCalculationResponse{
		Id: event.Id,
	}, nil
}

func NewCalculationService(config CalculationServiceConfig) CalculationService {
	return &calculationServiceImpl{
		client:       eventbridge.NewFromConfig(config.Aws),
		eventBusName: config.EventBusName,
	}
}
