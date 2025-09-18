package repositories

import (
	"editor/app/domain"
	"editor/app/types"
	"editor/app/utils"
	"encoding/json"
	"fmt"
	"os"
)

type messageRepository struct {
	dataDir string
	path    string
}

func NewMessageRepo(dataDir string) types.MessageRepository {
	repo := messageRepository{
		dataDir: dataDir,
		path:    fmt.Sprintf("%s/messages.json", dataDir),
	}
	return &repo
}

func (r *messageRepository) RestoreAll() (messages types.Messages, err error) {
	if !utils.CheckFile(r.path) {
		return domain.InitMessages([]types.Message{}), nil
	}
	data, err := os.ReadFile(r.path)
	if err != nil {
		return
	}
	messagesData := []types.MessageData{}
	err = json.Unmarshal(data, &messagesData)
	if err != nil {
		return
	}
	messages = domain.InitMessages(nil)
	for _, messageData := range messagesData {
		item := messages.FindById(messageData.MessageId)
		if item != nil {
			item.InitLanguage(messageData)
			continue
		}
		item = domain.InitMessage(messageData)
		messages.Add(item)
	}
	return messages, nil
}

func (r *messageRepository) PersistAll(messages types.Messages) (err error) {
	utils.CreateFileIfNotExist(r.path)
	data, err := json.Marshal(messages.ExportData())
	if err != nil {
		return
	}
	err = os.WriteFile(r.path, data, 0644)
	return
}
