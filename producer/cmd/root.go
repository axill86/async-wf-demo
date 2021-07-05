package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "producer",
	Short: "A producer for a async demo",
	Long: `This is the helper tool for submitting tasks for long-run async calculation.
Required env setup is the following:
aws profile should be set. You can use env variable AWS_PROFILE.
aws region should be set as well. Use AWS_REGION for that
eventbus name should be set. Use env variable MINICOMPUTE_EVENTBUS_NAME for that`,
}

//put initialize code here
func init() {
	rootCmd.AddCommand(getCreateCmd())
}
func Execute() {
	rootCmd.Execute()
}
