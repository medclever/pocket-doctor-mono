package types

// Message
type MessageData struct {
	MessageId    string `json:"message_id"`
	Message      string `json:"message"`
	LanguageCode string `json:"language_code"`
	DateCreate   string `json:"date_create"`
}
type Message interface {
	GetId() string
	InitLanguage(data MessageData)
	HasTranslate(languageCode string) bool
	AddLanguage(languageCode, text string)
	View() string
	ExportData() []MessageData
}
type Messages interface {
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
