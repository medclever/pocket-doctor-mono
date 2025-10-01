package domain

import (
	. "editor/app/types"
)

type articles struct {
}

func InitArticles(entyties []Article) Articles {
	return &articles{}
}

func (*articles) HasById(articleId ArticleId) bool {
	return false
}

func (*articles) FindById(articleId ArticleId) Article {
	return nil
}

func (*articles) Add(article Article) {

}

func (*articles) GetArticles() []Article {
	return nil
}

func (*articles) ExportData() []ArticleData {
	return nil
}
