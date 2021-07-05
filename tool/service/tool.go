package service

import "context"

type ToolServiceConfig struct {
}

type ToolRequest struct {
	RequestId string `json:"request_id"`
	Initiator string `json:"initiator"`
	Order     int    `json:"order"`
}

type ToolResponse struct {
	RequestId     string `json:"request_id"`
	StatusMessage string `json:"status_message"`
}

type Tool interface {
	ServeRequest(ctx context.Context, request ToolRequest) (*ToolResponse, error)
}

type toolImpl struct {
	config ToolServiceConfig
}

//ServeRequest just a dummy service for demo purpose
func (t *toolImpl) ServeRequest(ctx context.Context, request ToolRequest) (*ToolResponse, error) {
	return &ToolResponse{
		RequestId:     request.RequestId,
		StatusMessage: "Success",
	}, nil
}

func NewToolService(config ToolServiceConfig) Tool {
	return &toolImpl{config: config}
}
