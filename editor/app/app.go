package app

import (
	"editor/app/repositories"
	"editor/app/types"
)

type App struct {
	messageRepository types.MessageRepository
}

func New() *App {
	messageRepository := repositories.NewMessageRepo()
	app := App{
		messageRepository: messageRepository,
	}
	return &app
}

func (a *App) Init() {

}

func (a *App) MessagesGetList() []types.Message {
	return a.messageRepository.GetList(types.MessagesGetListParams{})
}
