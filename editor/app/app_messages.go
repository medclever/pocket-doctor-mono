package app

import (
	"editor/app/domain"
	"editor/app/errors"
	"editor/app/types"
	"fmt"
)

func (a *App) MessagesInit(messages ...types.Message) *App {
	for _, message := range messages {
		a.messages.Add(message)
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
		return fmt.Errorf("%w. code = %s", errors.LanguageNotFind, languageCode)
	}
	item := domain.CreateMessage(languageCode, message, a.timeNow.Now())
	a.messages.Add(item)
	return nil
}

func (a *App) MessagesTranslate(messageId, languageCode, message string) error {
	hasLanguage := a.languages.HasByCode(languageCode)
	if !hasLanguage {
		return fmt.Errorf("%w. code = %s", errors.LanguageNotFind, languageCode)
	}
	item := a.messages.FindById(messageId)
	if item == nil {
		return fmt.Errorf("%w. id = %s", errors.MessageNotFind, messageId)
	}
	hasTranslate := item.HasTranslate(languageCode)
	if hasTranslate {
		return fmt.Errorf("%w. code = %s", errors.MessageHasTranslate, languageCode)
	}
	item.AddTranslate(languageCode, message, a.timeNow.Now())
	return nil
}
