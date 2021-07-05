package config

import "os"

type ListenerConfig struct {
	InputQueue string
	StateMachine string
}

func ReadConfig() ListenerConfig {
	return ListenerConfig{
		InputQueue:   os.Getenv("COMPLETER_QUEUE"),
		StateMachine: os.Getenv("COMPLETER_SFN"),
	}
}

