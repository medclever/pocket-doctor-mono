package domain

import "editor/app/types"

type message struct {
	data types.MessageData
}

func NewMessage(data types.MessageData) types.Message {
	message := message{
		data: data,
	}
	return &message
}

func (m *message) View() string {
	return m.data.Message
}
