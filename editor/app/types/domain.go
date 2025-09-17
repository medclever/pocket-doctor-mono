package types

// Message
type MessageData struct {
	MessageId  string `json:"message_id"`
	Message    string `json:"message"`
	DateCreate string `json:"date_create"`
}
type Message interface {
	View() string
	ExportData() []MessageData
}
type Messages interface {
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
	View() string
	ExportData() []LanguageData
}
type Languages interface {
	Add(language Language)
	GetLanguages() []Language
	ExportData() []LanguageData
}
