package main

import (
	"strings"

	"github.com/memphisdev/memphis-functions.go/memphis"
)

func RemoveFields(message any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
	split_keys := strings.Split(inputs["keys"], ",")
	keys := make([]string, 0)
	for _, key := range split_keys {
		keys = append(keys, strings.TrimSpace(key))
	}

	RemoveFieldsInner := func(msgMapSubset *map[string]interface{}) {
		var RecursiveRemove func(*map[string]interface{})

		// Sorta like a depth first search but it deletes the keys specified on the way out
		RecursiveRemove = func(msgMapSubset *map[string]interface{}) {
			for _, value := range *msgMapSubset {
				if val_cast, ok := value.(map[string]interface{}); ok {
					RecursiveRemove(&val_cast)
				}
			}

			for _, value := range keys {
				delete(*msgMapSubset, value)
			}
		}
		RecursiveRemove(msgMapSubset)
	}

	event := message.(*map[string]any)

	RemoveFieldsInner(event)

	return event, headers, nil
}

func main() {
	var schema map[string]any
	memphis.CreateFunction(RemoveFields, memphis.PayloadAsJSON(&schema))
}
