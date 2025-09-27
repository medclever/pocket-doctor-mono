package app_test

import (
	"editor/app"
	"editor/app/errors"
	"editor/app/types"
	"editor/app/utils"
	"editor/app/utils/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

var date = "2025-01-01 00:00:00"
var timeNow = utils.NewTimeNowTest(date)
var lang = test.EnsureLang(date)
var message = test.EnsureMessage(date)

func Test_Messages_Translate_Success(t *testing.T) {
	assert := assert.New(t)
	app := app.New()
	app.LanguagesInit(lang("ru"), lang("en")).
		MessagesInit(message("mes1", "ru", "Сообщение 1")).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("mes1", "en", "Message 1")
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

func Test_Messages_Translate_Message_Not_Found(t *testing.T) {
	assert := assert.New(t)
	app := app.New()
	app.LanguagesInit(lang("ru"), lang("en")).
		MessagesInit(message("mes1", "ru", "Сообщение 1")).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("mes2", "en", "Message 1")
	assert.ErrorIs(err, errors.MessageNotFind)
}

func Test_Messages_Translate_Language_Not_Found(t *testing.T) {
	assert := assert.New(t)
	app := app.New()
	app.LanguagesInit(lang("ru")).
		MessagesInit(message("mes1", "ru", "Сообщение 1")).
		SetTimeNow(timeNow)

	err := app.MessagesTranslate("mes1", "en", "Message 1")
	assert.ErrorIs(err, errors.LanguageNotFind)
}

func Test_Messages_Translate_Has_Already_Translate(t *testing.T) {
	assert := assert.New(t)
	app := app.New().
		LanguagesInit(lang("ru"), lang("en")).
		SetTimeNow(timeNow)

	mes1 := message("mes1", "ru", "Сообщение 1")
	mes1.AddTranslate("en", "Message 1", timeNow.Now())
	app.MessagesInit(mes1)

	err := app.MessagesTranslate("mes1", "en", "Message 1")
	assert.ErrorIs(err, errors.MessageHasTranslate)
}
