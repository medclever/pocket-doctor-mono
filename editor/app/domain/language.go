package domain

import (
	"editor/app/types"
	"fmt"
	"time"
)

type language struct {
	data types.LanguageData
}

func InitLanguage(data types.LanguageData) types.Language {
	language := language{
		data: data,
	}
	return &language
}

func CreateLanguage(code, name string) types.Language {
	language := language{
		data: types.LanguageData{
			Code:       code,
			Name:       name,
			DateCreate: time.Now().Format(time.DateTime),
		},
	}
	return &language
}

func (m *language) View() string {
	text := fmt.Sprintf("code: %s, name: %s, date_create: %s\n", m.data.Code, m.data.Name, m.data.DateCreate)
	return text
}

func (m *language) ExportData() []types.LanguageData {
	return []types.LanguageData{m.data}
}
