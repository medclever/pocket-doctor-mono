package domain

import "editor/app/types"

type messages struct {
	messages []types.Message
}

func InitMessages(entyties []types.Message) types.Messages {
	return &messages{
		messages: entyties,
	}
}

func (m *messages) FindById(messageId string) types.Message {
	for _, message := range m.messages {
		if message.GetId() == messageId {
			return message
		}
	}
	return nil
}

func (m *messages) GetMessages() []types.Message {
	return m.messages
}

func (m *messages) Add(message types.Message) {
	m.messages = append(m.messages, message)
}

func (m *messages) ExportData() []types.MessageData {
	result := make([]types.MessageData, 0, len(m.messages))
	for _, message := range m.messages {
		result = append(result, message.ExportData()...)
	}
	return result
}
