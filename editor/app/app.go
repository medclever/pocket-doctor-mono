package app

import (
	"editor/app/domain"
	"editor/app/repositories"
	"editor/app/types"
	"editor/app/utils"
	lib_errors "errors"
	"fmt"
	"os"
)

type App struct {
	timeNow            types.TimeNow
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
		timeNow:            utils.NewTimeNowDefault(),
		messages:           domain.InitMessages([]types.Message{}),
		languages:          domain.InitLanguages([]types.Language{}),
		messageRepository:  messageRepository,
		languageRepository: languageRepository,
	}
	return &app
}

func (a *App) SetTimeNow(timeNow types.TimeNow) *App {
	a.timeNow = timeNow
	return a
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
		return lib_errors.Join(errs...)
	}
	return
}

func PrintErrorAndExit(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
