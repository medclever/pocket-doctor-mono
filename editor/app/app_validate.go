package app

import (
	"editor/app/errors"
	"fmt"
)

func (a *App) validateMessageAdd(languageCode string) error {
	return errors.Compose(
		a.validateLanguage(languageCode),
	)
}

func (a *App) validateMessageEdit(messageId, languageCode string) error {
	return errors.Compose(
		a.validateMessageId(messageId),
		a.validateLanguage(languageCode),
		a.validateMessageHasTranslate(messageId, languageCode),
	)
}

func (a *App) validateMessageTranslate(messageId, languageCode string) error {
	return errors.Compose(
		a.validateMessageId(messageId),
		a.validateLanguage(languageCode),
		a.validateMessageHasNotTranslate(messageId, languageCode),
	)
}

// ============

func (a *App) validateLanguage(languageCode string) func() error {
	return func() error {
		hasLanguage := a.languages.HasByCode(languageCode)
		if !hasLanguage {
			return fmt.Errorf("%w. code = %s", errors.LanguageNotFind, languageCode)
		}
		return nil
	}
}

func (a *App) validateMessageId(messageId string) func() error {
	return func() error {
		has := a.messages.HasById(messageId)
		if !has {
			return fmt.Errorf("%w. id = %s", errors.MessageNotFind, messageId)
		}
		return nil
	}
}

func (a *App) validateMessageHasNotTranslate(messageId, languageCode string) func() error {
	return func() error {
		item := a.messages.FindById(messageId)
		has := item.HasTranslate(languageCode)
		if has {
			return fmt.Errorf("%w. id = %s, lang = %s", errors.MessageHasTranslate, messageId, languageCode)
		}
		return nil
	}
}

func (a *App) validateMessageHasTranslate(messageId, languageCode string) func() error {
	return func() error {
		item := a.messages.FindById(messageId)
		has := item.HasTranslate(languageCode)
		if !has {
			return fmt.Errorf("%w. id = %s, lang = %s", errors.MessageHasNotTranslate, messageId, languageCode)
		}
		return nil
	}
}
