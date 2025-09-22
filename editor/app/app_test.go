package app_test

import (
	"editor/app"
	"editor/app/domain"
	"editor/app/errors"
	"editor/app/types"
	"editor/app/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var date = "2025-01-01 00:00:00"
var timeNow = utils.NewTimeNowTest(date)

func lang(code string) types.Language {
	return domain.InitLanguage(types.LanguageData{
		Code:       code,
		Name:       code,
		DateCreate: date,
	})
}
func message(id, languageCode, text string) types.Message {
	return domain.InitMessage(types.MessageData{
		MessageId:    id,
		Message:      text,
		LanguageCode: languageCode,
		DateCreate:   date,
	})
}

func Test_Messages_Translate_Success(t *testing.T) {
	assert := assert.New(t)
	app := app.New()
	app.LanguagesInit(lang("ru"), lang("en")).
		MessagesInit(message("mes1", "ru", "Сообщение 1")).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("en", "mes1", "Message 1")
	assert.Equal(app.MessagesGet("mes1").ExportData(), []types.MessageData{
		{
			MessageId:    "mes1",
			Message:      "Сообщение 1",
			LanguageCode: "ru",
			DateCreate:   date,
		},
		{
			MessageId:    "mes1",
			Message:      "Message 1",
			LanguageCode: "en",
			DateCreate:   date,
		},
	})
	assert.Nil(err)
}

func Test_Messages_Translate_Language_Not_Found(t *testing.T) {
	assert := assert.New(t)
	app := app.New()
	app.LanguagesInit(lang("ru")).
		MessagesInit(message("mes1", "ru", "Сообщение 1")).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("en", "mes1", "Message 1")
	assert.ErrorIs(err, errors.NotFindLanguage)
}

func Test_Messages_Translate_Has_Already_Translate(t *testing.T) {
	assert := assert.New(t)
	app := app.New()

	mes1 := message("mes1", "ru", "Сообщение 1")
	mes1.AddLanguage("en", "Message 1", timeNow.Now())
	app.LanguagesInit(lang("ru"), lang("en")).
		MessagesInit(mes1).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("en", "mes1", "Message 1")
	assert.ErrorIs(err, errors.MessageHasTranslate)
}
