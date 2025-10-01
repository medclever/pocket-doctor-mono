package app

import (
	"editor/app/types"
)

func (a *App) ArticlesInit(articles ...types.Article) types.App {
	for _, article := range articles {
		a.articles.Add(article)
	}
	return a
}

func (a *App) ArticlesGetList() types.Articles {
	return a.articles
}

func (a *App) ArticlesAdd(languageCode string, afterId *types.ArticleId, title string) {

}
