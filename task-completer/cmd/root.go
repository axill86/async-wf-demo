package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "task-completer",
	Short: "task-completer app which completes the step function execution",
	Long:  `This is a task-completer tool for demo purpose.
It supposed to listen for COMPLETER_QUEUE and complete tasks in state machine.
Also lambda is available.
Following env variables need to be set
AWS_PROFILE - aws profile to be used (need to be set up)
AWS_REGION - aws region to be used
COMPLETER_QUEUE - url of input sqs queue
`,
}
func init() {
	rootCmd.AddCommand(getListenCommand())
}

func Execute() {
	rootCmd.Execute()
}