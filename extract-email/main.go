package main

import (
	"github.com/memphisdev/memphis-functions.go/memphis"
	"regexp"
)


type ConversionError struct {
	message string
}

func (e *ConversionError) Error() string {
	return e.message
}

type NoEmailsError struct {
	message string
}

func (e *NoEmailsError) Error() string {
	return e.message
}

var email_regex string
var re *regexp.Regexp

func EventHandler(message any, headers map[string]string, inputs map[string]string) (any, map[string]string,  error){
	event := *message.(*map[string]any)

	strWithEmail, ok := event[inputs["email"]].(string);
	if !ok{
		return nil, nil, &ConversionError{message: "The given event[inputs['email']] field was not of type string."}
	}

	emails := re.FindAllString(strWithEmail, -1)
	
	if emails != nil{
		event[inputs["out"]] = emails
	}else{
		return nil, nil, &NoEmailsError{message: "There were no emails found in this event"} 
	}

	return event, headers, nil
}

func main() {	
	email_regex = `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
	re = regexp.MustCompile(email_regex)
	
	var schema map[string]any
	memphis.CreateFunction(EventHandler, memphis.PayloadAsJSON(&schema))
}