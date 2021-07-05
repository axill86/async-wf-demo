package endpoint

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
	"task-completer/service"
	"time"
)

type QueueListenerConfig struct {
	Aws                 aws.Config
	Queue               string
	Interval            time.Duration
	OnMessage           func(request service.CompleteTaskRequest) error
}

type QueueListener interface {
	Start(interrupt <-chan bool)
}

type queueListenerImpl struct {
	client              *sqs.Client
	interval            time.Duration
	queueUrl            string
	onMessage           func(request service.CompleteTaskRequest) error
}

func (q *queueListenerImpl) Start(interrupt <-chan bool) {
	t := time.NewTicker(q.interval)
	go func() {
		for {
			select {
			case <-t.C:
				log.Println("Reading messages")
				output, err := q.client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
					QueueUrl:            aws.String(q.queueUrl),
				})
				if err != nil {
					log.Panicf("failed to read messages %e", err)
				}
				for _, message := range output.Messages {
					//unmarshall request
					var request service.CompleteTaskRequest
					if err := json.Unmarshal([]byte(*message.Body), &request); err != nil {
						log.Panicf("failed to parse message %e", err)
					}
					//trigger callback
					if err := q.onMessage(request); err != nil {
						log.Panicf("failed to execute callback %e", err)
					}
					//delete message if processed
					q.client.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
						QueueUrl:      aws.String(q.queueUrl),
						ReceiptHandle: message.ReceiptHandle,
					})
				}
			case <-interrupt:
				return
			}
		}
	}()
}

func NewQueueListener(config *QueueListenerConfig) QueueListener {
	client := sqs.NewFromConfig(config.Aws)
	return &queueListenerImpl{
		client:              client,
		interval:            config.Interval,
		queueUrl:            config.Queue,
		onMessage:           config.OnMessage,
	}
}
