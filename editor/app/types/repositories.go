package types

type MessageRepository interface {
	RestoreAll() (messages []Message, err error)
	PersistAll(messages []Message) error
}

type MessagesGetListParams struct{}
