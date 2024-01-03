package main

import (
	"strconv"

	"github.com/memphisdev/memphis-functions.go/memphis"
)

type ConversionError struct {
	message string
}

func (e *ConversionError) Error() string {
	return e.message
}

func CheckSeverity(message any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
	event := *message.(*map[string]any)

	var measuredValue float64

	if msgField, ok := event[inputs["field"]].(float64); ok {
		measuredValue = msgField
	} else {
		return nil, nil, &ConversionError{message: "Given field key was not able to be converted to a float"}
	}

	severityCutoff, err := strconv.ParseFloat(inputs["cutoff"], 32)

	if err != nil {
		return nil, nil, err
	}

	if measuredValue >= severityCutoff {
		(event)["severity"] = inputs["high"]
	} else {
		(event)["severity"] = inputs["low"]
	}

	return event, headers, nil
}

func main() {
	var schema map[string]any
	memphis.CreateFunction(CheckSeverity, memphis.PayloadAsJSON(&schema))
}
