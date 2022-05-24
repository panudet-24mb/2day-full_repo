package services

import (
	"account_gateway/events"
	"account_gateway/internal/commands"
	"errors"
	"log"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) (id string, err error)
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer}
}

func (obj accountService) OpenAccount(command commands.OpenAccountCommand) (id string, err error) {

	if command.AccountName == "" || command.AccountType == 0 || command.AccountNumber == "" {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		AccountType:   command.AccountType,
		AccountNumber: command.AccountNumber,
		AccountName:   command.AccountName,
	}

	log.Printf("%#v", event)
	return event.AccountNumber, obj.eventProducer.Produce(event)
}
