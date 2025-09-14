package domain

import (
	"editor/app/types"
	"editor/app/utils"
	"fmt"
	"time"
)

type message struct {
	data types.MessageData
}

func NewMessage(data types.MessageData) types.Message {
	message := message{
		data: data,
	}
	return &message
}

func NewMessageFromText(text string) types.Message {
	message := message{
		data: types.MessageData{
			MessageId:  utils.GenereateUid(10),
			Message:    text,
			DateCreate: time.Now().Format(time.DateTime),
		},
	}
	return &message
}

func (m *message) View() string {
	text := fmt.Sprintf("id: %s date_create: %s\n", m.data.MessageId, m.data.DateCreate)
	text += m.data.Message
	return text
}

func (m *message) ExportData() types.MessageData {
	return m.data
}
