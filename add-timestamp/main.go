package main

import (
	"github.com/memphisdev/memphis-functions.go/memphis"
	"time"
)

func AddTimestamp(message any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
	// Assumes JSON encoding
	event := *message.(*map[string]any)

	event["timestamp"] = time.Now().Round(time.Second)

	return event, headers, nil
}

func main() {
	var schema map[string]any
	memphis.CreateFunction(AddTimestamp, memphis.PayloadAsJSON(&schema))
}
