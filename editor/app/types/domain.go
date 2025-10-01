package types

import (
	"encoding/json"
	"time"
)

type ArticleData struct {
	ArticleId  ArticleId         `json:"artice_id"`
	Title      MessageParams     `json:"title"`
	Necessary  []ArticleItemData `json:"necessary"`
	Possible   []ArticleItemData `json:"possible"`
	MustNot    []ArticleItemData `json:"must_not"`
	Important  []ArticleItemData `json:"important"`
	Text       []ArticleItemData `json:"text"`
	DateCreate string            `json:"date_create"`
}
type ArticleItemData struct {
	ItemId   ItemId          `json:"item_id"`
	ItemType ItemType        `json:"item_type"`
	Params   json.RawMessage `json:"params"`
}
type MessageParams struct {
	MessageId string `json:"message_id"`
}
type GalleryParams struct {
	ImagesIds []string `json:"image_ids"`
}
type LinkParams struct {
	ArticleId string `json:"article_id"`
}
type Article interface {
	GetId() ArticleId
	AddItem(afterId *ItemId, place PlaceType, itemType ItemType, params any)
	EditItem(itemId ItemId, params any)
	DeleteItem(itemId ItemId)
	View(languageCode string) string
	ExportData() []ArticleData
}
type ArticleItem interface {
	GetId() ItemId
	Edit(params any)
	View(languageCode string) string
	ExportData() []ArticleItemData
}
type Articles interface {
	HasById(articleId ArticleId) bool
	FindById(articleId ArticleId) Article
	Add(article Article)
	GetArticles() []Article
	ExportData() []ArticleData
}

// Message
type MessageData struct {
	MessageId    string `json:"message_id"`
	Message      string `json:"message"`
	LanguageCode string `json:"language_code"`
	DateCreate   string `json:"date_create"`
}
type Message interface {
	GetId() string
	Edit(languageCode, text string)
	GetTranslate(languageCode string) *MessageData
	InitTranslate(data MessageData)
	AddTranslate(languageCode, text string, now time.Time)
	HasTranslate(languageCode string) bool
	View(languageCode string) string
	Meta(languageCode string) string
	ExportData() []MessageData
}
type Messages interface {
	HasById(messageId string) bool
	FindById(messageId string) Message
	Add(message Message)
	GetMessages() []Message
	ExportData() []MessageData
}

// Language
type LanguageData struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	DateCreate string `json:"date_create"`
}
type Language interface {
	GetCode() string
	View() string
	ExportData() []LanguageData
}
type Languages interface {
	HasByCode(languageCode string) bool
	FindByCode(languageCode string) Language
	Add(language Language)
	GetLanguages() []Language
	ExportData() []LanguageData
}
