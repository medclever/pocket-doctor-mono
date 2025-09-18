package domain

import "editor/app/types"

type languages struct {
	languages []types.Language
}

func InitLanguages(entyties []types.Language) types.Languages {
	return &languages{
		languages: entyties,
	}
}

func (m *languages) HasByCode(languageCode string) bool {
	return m.FindByCode(languageCode) != nil
}

func (m *languages) FindByCode(languageCode string) types.Language {
	for _, language := range m.languages {
		if language.GetCode() == languageCode {
			return language
		}
	}
	return nil
}

func (m *languages) GetLanguages() []types.Language {
	return m.languages
}

func (m *languages) Add(language types.Language) {
	m.languages = append(m.languages, language)
}

func (m *languages) ExportData() []types.LanguageData {
	result := make([]types.LanguageData, 0, len(m.languages))
	for _, language := range m.languages {
		result = append(result, language.ExportData()...)
	}
	return result
}
