package config

import "os"

type SQS struct {
	//Required param
	OutputQueue string
	//Optional. for commandline usage only
	InputQueue string
}

func ReadSQSConfig() *SQS {
	return &SQS{
		OutputQueue: os.Getenv("TOOL_OUTPUT_QUEUE_URL"),
		InputQueue:  os.Getenv("TOOL_INPUT_QUEUE_URL"),
	}
}
