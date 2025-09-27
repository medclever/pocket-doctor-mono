package errors

import "errors"

var (
	MessageNotFind         = errors.New("Message not find")
	MessageHasTranslate    = errors.New("Message already has translate")
	MessageHasNotTranslate = errors.New("Message has`t translate")
	LanguageNotFind        = errors.New("Language not find")
)

func Compose(fns ...func() error) error {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}
