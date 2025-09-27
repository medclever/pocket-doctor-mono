package types

type App interface {
	SetTimeNow(timeNow TimeNow) App
	Init()
	Persist() (err error)
	//
	LanguagesInit(languages ...Language) App
	LanguagesGetList() Languages
	LanguagesAdd(code, name string)
	//
	MessagesInit(messages ...Message) App
	MessagesGetList() Messages
	MessagesGet(messageId string) Message
	MessagesAdd(languageCode, message string) error
	MessagesEdit(messageId, languageCode, message string) error
	MessagesTranslate(messageId, languageCode, message string) error
}
