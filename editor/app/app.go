package app

import (
	"editor/app/domain"
	"editor/app/errors"
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

func (a *App) MessagesInit(messages ...types.Message) *App {
	for _, message := range messages {
		a.messages.Add(message)
	}
	return a
}

func (a *App) LanguagesInit(languages ...types.Language) *App {
	for _, language := range languages {
		a.languages.Add(language)
	}
	return a
}

func (a *App) MessagesGetList() types.Messages {
	return a.messages
}

func (a *App) MessagesGet(messageId string) types.Message {
	return a.messages.FindById(messageId)
}

func (a *App) MessagesAdd(languageCode, message string) error {
	hasLanguage := a.languages.HasByCode(languageCode)
	if !hasLanguage {
		return fmt.Errorf("%w. code = %s", errors.NotFindLanguage, languageCode)
	}
	item := domain.CreateMessage(languageCode, message, a.timeNow.Now())
	a.messages.Add(item)
	return nil
}

func (a *App) MessagesTranslate(languageCode, messageId, message string) error {
	hasLanguage := a.languages.HasByCode(languageCode)
	if !hasLanguage {
		return fmt.Errorf("%w. code = %s", errors.NotFindLanguage, languageCode)
	}
	item := a.messages.FindById(messageId)
	if item == nil {
		a.MessagesAdd(languageCode, message)
		return nil
	}
	hasTranslate := item.HasTranslate(languageCode)
	if hasTranslate {
		return fmt.Errorf("%w. code = %s", errors.MessageHasTranslate, languageCode)
	}
	item.AddLanguage(languageCode, message, a.timeNow.Now())
	return nil
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
