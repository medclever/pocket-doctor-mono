package domain

import (
	"editor/app/types"
	"editor/app/utils"
	"fmt"
	"time"
)

type message struct {
	id    string
	items []types.MessageData
}

func InitMessage(data types.MessageData) types.Message {
	message := message{
		id:    data.MessageId,
		items: []types.MessageData{data},
	}
	return &message
}

func CreateMessage(languageCode, text string) types.Message {
	message := message{
		items: []types.MessageData{
			createMessage(languageCode, text),
		},
	}
	return &message
}

func (m *message) GetId() string {
	return m.id
}

func (m *message) InitLanguage(data types.MessageData) {
	m.items = append(m.items, data)
}

func (m *message) AddLanguage(languageCode, text string) {
	m.items = append(m.items, createMessage(languageCode, text))
}

func (m *message) HasTranslate(languageCode string) bool {
	for _, item := range m.items {
		if item.LanguageCode == languageCode {
			return true
		}
	}
	return false
}

func (m *message) View() string {
	text := ""
	fmt.Println(m.items)
	for _, item := range m.items {
		text = fmt.Sprintf("id: %s, language_code: %s, date_create: %s\n",
			item.MessageId,
			item.LanguageCode,
			item.DateCreate,
		)
		text += item.Message + "\n"
	}
	return text
}

func (m *message) ExportData() []types.MessageData {
	return m.items
}

func createMessage(languageCode, text string) types.MessageData {
	return types.MessageData{
		MessageId:    utils.GenereateUid(10),
		Message:      text,
		LanguageCode: languageCode,
		DateCreate:   time.Now().Format(time.DateTime),
	}
}
