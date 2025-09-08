package domain

import "editor/app/types"

type message struct {
	message string
}

func NewMessage(text string) types.Message {
	message := message{
		message: text,
	}
	return &message
}

func (m *message) View() string {
	return m.message
}
