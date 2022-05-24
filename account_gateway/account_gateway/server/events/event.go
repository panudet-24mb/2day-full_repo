package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
}

type Event interface {
}

type OpenAccountEvent struct {
	AccountType   int
	AccountNumber string
	AccountName   string
}
