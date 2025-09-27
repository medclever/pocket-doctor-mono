package app

import (
	"editor/app/domain"
	"editor/app/types"
)

func (a *App) LanguagesInit(languages ...types.Language) *App {
	for _, language := range languages {
		a.languages.Add(language)
	}
	return a
}

func (a *App) LanguagesGetList() types.Languages {
	return a.languages
}

func (a *App) LanguagesAdd(code, name string) {
	a.languages.Add(domain.CreateLanguage(code, name))
}
