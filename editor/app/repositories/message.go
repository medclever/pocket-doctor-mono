package repositories

import (
	"editor/app/domain"
	"editor/app/types"
	"encoding/json"
	"os"
)

type messageRepository struct {
	file string
}

func NewMessageRepo(file string) types.MessageRepository {
	repo := messageRepository{
		file: file,
	}
	return &repo
}

func (r *messageRepository) RestoreAll() (messages []types.Message, err error) {
	data, err := os.ReadFile(r.file)
	if err != nil {
		return
	}
	messagesData := []types.MessageData{}
	err = json.Unmarshal(data, &messagesData)
	if err != nil {
		return
	}
	for _, messageData := range messagesData {
		messages = append(messages, domain.NewMessage(messageData))
	}
	return
}

func (r *messageRepository) PersistAll(messages []types.Message) error {
	return nil
}
