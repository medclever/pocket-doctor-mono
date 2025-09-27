package app

import (
	"editor/app/domain"
	"editor/app/types"
)

func (a *App) MessagesInit(messages ...types.Message) types.App {
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
	if err := a.validateMessageAdd(languageCode); err != nil {
		return err
	}
	item := domain.CreateMessage(languageCode, message, a.timeNow.Now())
	a.messages.Add(item)
	return nil
}

func (a *App) MessagesEdit(messageId, languageCode, message string) error {
	if err := a.validateMessageEdit(messageId, languageCode); err != nil {
		return err
	}
	item := a.messages.FindById(messageId)
	item.Edit(languageCode, message)
	return nil
}

func (a *App) MessagesTranslate(messageId, languageCode, message string) error {
	if err := a.validateMessageTranslate(messageId, languageCode); err != nil {
		return err
	}
	item := a.messages.FindById(messageId)
	item.AddTranslate(languageCode, message, a.timeNow.Now())
	return nil
}
