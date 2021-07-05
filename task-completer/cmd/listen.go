package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	appconfig "task-completer/config"
	"task-completer/endpoint"
	"task-completer/service"
	"time"
)

func getListenCommand() *cobra.Command {
	conf := appconfig.ReadConfig()
	var interval time.Duration
	var inputQueueUrl string
	
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "listen to queue",
		Long:  `Listenes for queue and completes state machine`,
		Run: func(cmd *cobra.Command, args []string) {
			awsConfig, err := config.LoadDefaultConfig(context.Background())
			if err != nil {
				log.Panicf("failed to load default config %e", err)
			}
			completerService := service.NewCompleter(service.CompleterConfig{
				AWS:          awsConfig,
			})


			listener := endpoint.NewQueueListener(&endpoint.QueueListenerConfig{
				Aws:      awsConfig,
				Queue:    inputQueueUrl,
				Interval: interval,
				OnMessage: func(request service.CompleteTaskRequest) error {
					ctxt := context.Background()
					log.Printf("Received request %v", request)
					if err := completerService.CompleteTask(ctxt, request); err != nil {
						return err
					}
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
	return cmd
}
