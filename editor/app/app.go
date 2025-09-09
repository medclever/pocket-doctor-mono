package app

import (
	"editor/app/repositories"
	"editor/app/types"
	"fmt"
	"os"
)

type App struct {
	messageRepository types.MessageRepository
	messages          []types.Message
}

func New() *App {
	messageRepository := repositories.NewMessageRepo("./data/messages.ru.json")
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

func (a *App) MessagesGetList() []types.Message {
	return a.messages
}

func PrintErrorAndExit(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
