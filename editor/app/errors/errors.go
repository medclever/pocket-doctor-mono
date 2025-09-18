package errors

import "errors"

var (
	NotFindLanguage     = errors.New("Language not find")
	MessageHasTranslate = errors.New("Message already has translate")
)
