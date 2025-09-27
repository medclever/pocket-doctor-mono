package errors

import "errors"

var (
	MessageNotFind      = errors.New("Message not find")
	MessageHasTranslate = errors.New("Message already has translate")
	LanguageNotFind     = errors.New("Language not find")
)
