package types

type MessageRepository interface {
	RestoreAll() (messages Messages, err error)
	PersistAll(messages Messages) error
}

type LanguageRepository interface {
	RestoreAll() (languages Languages, err error)
	PersistAll(languages Languages) error
}
