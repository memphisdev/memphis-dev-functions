package main

import (
	"github.com/memphisdev/memphis-functions.go/memphis"
	"net/http"
	"io"
	"fmt"
)

type ConversionError struct {
	message string
}

func (e *ConversionError) Error() string {
	return e.message
}

func EventHandler(message any, headers map[string]string, inputs map[string]string) (any, map[string]string,  error){
	event := *message.(*map[string]any)

	ip, ok := event[inputs["geolocation"]].(string)

	if !ok{
		return nil, nil, &ConversionError{message: "The event field geolocation points to could not be cast to a string"}
	}

	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	res, err := http.Get(url)

	if err != nil{
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err != nil{
		return nil, nil, err
	}

	event[inputs["out"]] = string(body) 
	
	return event, headers, nil
}

func main() {	
	var schema map[string]any
	memphis.CreateFunction(EventHandler, memphis.PayloadAsJSON(&schema))
}