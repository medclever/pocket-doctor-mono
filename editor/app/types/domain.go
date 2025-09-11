package types

type MessageData struct {
	MessageId  string `json:"message_id"`
	Message    string `json:"message"`
	DateCreate string `json:"date_create"`
}

type Message interface {
	View() string
	ExportData() MessageData
}

type Messages interface {
	Add(message Message)
	GetMessages() []Message
	ExportData() []MessageData
}
