package types

type ArticleRepository interface {
	RestoreAll() (articles Articles, err error)
	PersistAll(articles Articles) error
}

type MessageRepository interface {
	RestoreAll() (messages Messages, err error)
	PersistAll(messages Messages) error
}

type LanguageRepository interface {
	RestoreAll() (languages Languages, err error)
	PersistAll(languages Languages) error
}
