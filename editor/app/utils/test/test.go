package test

import (
	"editor/app/domain"
	"editor/app/types"
)

func EnsureLang(date string) func(code string) types.Language {
	return func(code string) types.Language {
		return domain.InitLanguage(types.LanguageData{
			Code:       code,
			Name:       code,
			DateCreate: date,
		})
	}
}

func EnsureMessage(date string) func(id, languageCode, text string) types.Message {
	return func(id, languageCode, text string) types.Message {
		return domain.InitMessage(types.MessageData{
			MessageId:    id,
			Message:      text,
			LanguageCode: languageCode,
			DateCreate:   date,
		})
	}
}
