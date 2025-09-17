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
	messages           types.Messages
	messageRepository  types.MessageRepository
	languages          types.Languages
	languageRepository types.LanguageRepository
}

func New() *App {
	rootPath := "./data"
	messageRepository := repositories.NewMessageRepo(rootPath)
	languageRepository := repositories.NewLanguageRepo(rootPath)
	app := App{
		messageRepository:  messageRepository,
		languageRepository: languageRepository,
	}
	return &app
}

func (a *App) Init() {
	messages, err := a.messageRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.messages = messages

	languages, err := a.languageRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.languages = languages
}

func (a *App) Persist() (err error) {
	errs := []error{}

	err = a.messageRepository.PersistAll(a.messages)
	errs = append(errs, err)

	err = a.languageRepository.PersistAll(a.languages)
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
	a.messages.Add(domain.CreateMessage(message))
}

func (a *App) LanguagesGetList() types.Languages {
	return a.languages
}

func (a *App) LanguagesAdd(code, name string) {
	a.languages.Add(domain.CreateLanguage(code, name))
}

func PrintErrorAndExit(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
