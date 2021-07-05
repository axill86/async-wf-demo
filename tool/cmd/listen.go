package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	appconfig "tool/config"
	"tool/endpoint"
	"tool/service"
)

func getListenCommand() *cobra.Command {
	conf := appconfig.ReadSQSConfig()
	var interval time.Duration
	var inputQueueUrl string
	var outputQueueUrl string
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "starts listening for the queue",
		Long:  `Polls the input queue once per number of seconds until not interrupted`,
		Run: func(cmd *cobra.Command, args []string) {
			awsConfig, err := config.LoadDefaultConfig(context.Background())
			if err != nil {
				log.Panicf("failed to load default config %e", err)
			}
			reportService := service.NewSQSReporterService(service.SQSReporterServiceConfig{
				QueueUrl: outputQueueUrl,
				Aws:      awsConfig,
			})

			toolService := service.NewToolService(service.ToolServiceConfig{})
			listener := endpoint.NewQueueListener(&endpoint.QueueListenerConfig{
				Aws:      awsConfig,
				Queue:    inputQueueUrl,
				Interval: interval,
				OnMessage: func(request service.ToolRequest) error {
					ctxt := context.Background()
					log.Printf("Received request %v", request)
					resp, err := toolService.ServeRequest(ctxt, request)
					if err != nil {
						return err
					}
					sendResult, err := reportService.SendEvent(context.Background(), service.SQSReporterRequest{ToolResponse: *resp})
					if err != nil {
						return err
					}
					log.Printf("Sent message with result %v", *sendResult)
					return nil
				},
			})
			interrupt := make(chan bool)
			sigs := make(chan os.Signal)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			listener.Start(interrupt)
			<-sigs
			interrupt <- true
		},
	}
	cmd.Flags().DurationVarP(&interval, "interval", "i", 5*time.Second, "Poll interval")
	cmd.Flags().StringVar(&inputQueueUrl, "inputQueue", conf.InputQueue, "Input queue url")
	cmd.Flags().StringVar(&outputQueueUrl, "outputQueue", conf.OutputQueue, "Output queue url")
	return cmd
}
