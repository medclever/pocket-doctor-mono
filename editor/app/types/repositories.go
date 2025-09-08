package types

type MessageRepository interface {
	GetList(params MessagesGetListParams) []Message
}

type MessagesGetListParams struct{}
