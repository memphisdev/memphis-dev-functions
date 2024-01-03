package main

import (
	"github.com/memphisdev/memphis-functions.go/memphis"
)

// https://github.com/memphisdev/memphis.go#creating-a-memphis-function
func EventHandler(message any, headers map[string]string, inputs map[string]string) (any, map[string]string,  error){
	// Here is a short example of converting the message payload to bytes and back
	event := *message.(*map[string]any)
	event[inputs["field_to_ingest"]] = "Hello from Memphis!"
	
	// Return the payload back 
	return event, headers, nil
}


func main() {
	var schema map[string]any	
	memphis.CreateFunction(EventHandler, memphis.PayloadAsJSON(&schema))
}