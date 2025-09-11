package app

import (
	"editor/app/domain"
	"editor/app/repositories"
	"editor/app/types"
	"errors"
	"fmt"
	"os"
)

type App struct {
	messageRepository types.MessageRepository
	messages          types.Messages
}

func New() *App {
	messageRepository := repositories.NewMessageRepo("./data")
	app := App{
		messageRepository: messageRepository,
	}
	return &app
}

func (a *App) Init() {
	messages, err := a.messageRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.messages = messages
}

func (a *App) Persist() (err error) {
	errs := []error{}

	err = a.messageRepository.PersistAll(a.messages)
	errs = append(errs, err)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return
}

func (a *App) MessagesGetList() types.Messages {
	return a.messages
}

func (a *App) MessagesAdd(message string) {
	a.messages.Add(domain.NewMessageFromText(message))
}

func PrintErrorAndExit(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
