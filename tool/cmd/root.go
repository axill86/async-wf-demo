package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "tool is used for small minicompute demo",
	Long: `This is a helper tool for demo purpose.
It supposed to listen for TOOL_INPUT_QUEUE_URL and provide output to TOOL_OUTPUT_QUEUE_URL.
Also lambda is available.
Following env variables need to be set
AWS_PROFILE - aws profile to be used (need to be set up)
AWS_REGION - aws region to be used
TOOL_INPUT_QUEUE_URL - url of input sqs queue
TOOL_OUTPUT_QUEUE_URL - url of output sqs queue`,
}

func init() {
	rootCmd.AddCommand(getListenCommand())
}

func Execute() {
	rootCmd.Execute()
}
