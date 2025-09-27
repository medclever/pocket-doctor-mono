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

func CreateMessage(languageCode, text string, now time.Time) types.Message {
	message := message{
		items: []types.MessageData{
			createMessage(languageCode, text, now),
		},
	}
	return &message
}

func (m *message) GetId() string {
	return m.id
}

func (m *message) InitTranslate(data types.MessageData) {
	m.items = append(m.items, data)
}

func (m *message) AddTranslate(languageCode, text string, now time.Time) {
	if !m.HasTranslate(languageCode) {
		item := createMessage(languageCode, text, now)
		item.MessageId = m.id
		m.items = append(m.items, item)
	}
}

func (m *message) Edit(languageCode, text string) {
	item := m.GetTranslate(languageCode)
	if item != nil {
		item.Message = text
	}
}

func (m *message) GetTranslate(languageCode string) *types.MessageData {
	for i, item := range m.items {
		if item.LanguageCode == languageCode {
			return &m.items[i]
		}
	}
	return nil
}

func (m *message) HasTranslate(languageCode string) bool {
	return m.GetTranslate(languageCode) != nil
}

func (m *message) View(languageCode string) string {
	text := ""
	item := m.GetTranslate(languageCode)
	if item != nil {
		text = item.Message
	}
	return text
}

func (m *message) Meta(languageCode string) string {
	text := ""
	item := m.GetTranslate(languageCode)
	if item != nil {
		text = fmt.Sprintf("id: %s, lang_code: %s, date_create: %s", item.MessageId, item.LanguageCode, item.DateCreate)
	}
	return text
}

func (m *message) ExportData() []types.MessageData {
	return m.items
}

func createMessage(languageCode, text string, now time.Time) types.MessageData {
	return types.MessageData{
		MessageId:    utils.GenereateUid(10),
		Message:      text,
		LanguageCode: languageCode,
		DateCreate:   now.Format(time.DateTime),
	}
}
