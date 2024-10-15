package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MyJson struct {
	Mystery     any    `json:"mystery"`
	StringField string `json:"string_field"`
}

func main() {
	var jsonParsed MyJson
	json.Unmarshal([]byte(`{"mystery": {"innerArray": [1,2,3]}, "string_field": "this is a string"}`), &jsonParsed)

	switch v := jsonParsed.Mystery.(type) {
	case map[string]any:
		val, ok := v["innerArray"]
		if ok {
			switch v2 := val.(type) {
			case []any:
				for _, element := range v2 {
					switch element.(type) {
					case int:
						fmt.Printf("Found int: %d\n", element)
					// case float64:
					// 	fmt.Printf("Found float64: %f\n", element)
					default:
						fmt.Printf("Unknown Type: %v\n", reflect.TypeOf(element))
						fmt.Printf("Value: %v\n", element)
					}
				}

			default:
				fmt.Printf("Unknown innerArray type: %v", reflect.TypeOf(v2))
			}
			return
		}
		fmt.Println("innerArray not exist")
	default:
		fmt.Printf("Unknown type: %v", reflect.TypeOf(jsonParsed))
	}
}
