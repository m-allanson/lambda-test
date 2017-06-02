package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

const myInput = "hellooo CI world"

// Echo returns its input
func Echo(input string) string {
	return input
}

// Handle runs every time lambda is invoked
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	result := Echo(myInput)
	log.Printf("Handle result is: %s\n", result)
	return nil, nil
}

func main() {
	result := Echo(myInput)
	log.Printf("main result is: %s\n", result)
}
