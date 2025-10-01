package domain

import (
	. "editor/app/types"
)

type article struct {
}

func InitArticle(data ArticleData) Article {
	return &article{}
}

func (a *article) GetId() ArticleId {
	return ArticleId("")
}

func (a *article) AddItem(
	afterId *ItemId,
	place PlaceType,
	itemType ItemType,
	params any,
) {

}
func (a *article) EditItem(itemId ItemId, params any) {

}

func (a *article) DeleteItem(itemId ItemId) {

}

func (a *article) View(languageCode string) string {
	return ""
}

func (a *article) ExportData() []ArticleData {
	return nil
}
