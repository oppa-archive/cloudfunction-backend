package models

import (
	_ "cloud.google.com/go/firestore"
)

type Greeting struct {
	Name    string `dynamodbav:"name" json:"name"`
	Message string `dynamodbav:"message" json:"message"`
}

type GreetingRepository interface {
	AddGreeting(greeting *Greeting) (*Greeting, error)
	GetGreeter(name string) (*Greeting, error)
}
