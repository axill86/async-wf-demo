package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"log"
	appconfig "sandbox/producer/config"
	"sandbox/producer/service"
	"strconv"
)

func getCreateCmd() *cobra.Command {

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "creates a task for long-running job",
		Long: `Creates a task for long-running calculation process.
Returns the id of created task`,
		Args: nil,
		Run: func(cmd *cobra.Command, args []string) {
			aws, err := config.LoadDefaultConfig(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			appConfig := appconfig.ReadConfig()
			svc := service.NewCalculationService(service.CalculationServiceConfig{
				Aws:          aws,
				EventBusName: appConfig.EventBusName,
			})
			order, _ := strconv.Atoi(cmd.Flag("order").Value.String())
			request := service.CreateCalculationRequest{
				Initiator: cmd.Flag("initiator").Value.String(),
				Order:     order,
				Flow:      cmd.Flag("flow").Value.String(),
			}
			resp, err := svc.CreateCalculation(context.Background(), request)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v", resp)
		},
	}
	createCmd.Flags().String("initiator", "", "initiator of calculation")
	createCmd.MarkFlagRequired("initiator")
	createCmd.Flags().Int("order", -1, "order number")
	createCmd.MarkFlagRequired("order")
	createCmd.Flags().String("flow", service.FlowDummy, "Flow type for calculation")
	return createCmd
}
