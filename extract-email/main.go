package main

import (
	"encoding/json"
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

func EventHandler(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string,  error){
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil{
		return nil, nil, err
	}

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

	if eventBytes, err := json.Marshal(event); err == nil {
		return eventBytes, headers, nil	
	} else{
		return nil, nil, err
	}
}

func main() {	
	email_regex = `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
	re = regexp.MustCompile(email_regex)
	
	memphis.CreateFunction(EventHandler)
}