package main

import (
	"fmt"
	"github.com/memphisdev/memphis-functions.go/memphis"
)

func FlattenMessages(message any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
	flatten := func(out_map map[string]interface{}, value interface{}, parent_key string) {
		var recursiveFlatten func(map[string]interface{}, interface{}, string)

		recursiveFlatten = func(out_map map[string]interface{}, value interface{}, parent_key string) {
			switch value_typed := value.(type) {
			default:
				out_map[parent_key] = value
			case map[string]interface{}:
				if len(value_typed) == 0{
					out_map[parent_key] = value
					return
				}
				for key, value := range value_typed {
					var new_key string
					if len(parent_key) > 0{
						new_key = fmt.Sprintf("%s.%s", parent_key, key)
					}else{
						new_key = key
					}

					recursiveFlatten(out_map, value, new_key)
				}
			case []interface{}:
				if len(value_typed) == 0{
					out_map[parent_key] = value
					return
				}
				for index, value := range value_typed {
					new_key := fmt.Sprintf("%s_%d", parent_key, index)
					recursiveFlatten(out_map, value, new_key)
				}
			}
		}

		recursiveFlatten(out_map, value, parent_key)
	}

	out_struct := make(map[string]interface{})

	event := *message.(*map[string]any)

	flatten(out_struct, event, "")

	return out_struct, headers, nil
}

func main() {
	var schema map[string]any
	memphis.CreateFunction(FlattenMessages, memphis.PayloadAsJSON(&schema))
}
