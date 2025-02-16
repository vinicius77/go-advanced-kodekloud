package main

import (
	"fmt"
)

type APIError struct {
	// your code goes here Message string
	Message string
	Code    int
	Details map[string]interface{}
}

func (e *APIError) Error() string {
	// your code goes here
	return e.Message
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
	// your code goes here
	return &APIError{
		Message: message,
		Code:    code,
		Details: details,
	}
}

func main() {
	err := NewAPIError(400, "Bad request", map[string]interface{}{
		"field": "username",
		"error": "cannot be empty",
	})
	fmt.Println(err)
}
