package types

type MessageRepository interface {
	RestoreAll() (messages Messages, err error)
	PersistAll(messages Messages) error
}

type MessagesGetListParams struct{}
