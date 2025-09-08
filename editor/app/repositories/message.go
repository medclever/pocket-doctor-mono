package repositories

import (
	"editor/app/domain"
	"editor/app/types"
)

type messageRepository struct {
}

func NewMessageRepo() types.MessageRepository {
	repo := messageRepository{}
	return &repo
}

func (r *messageRepository) GetList(params types.MessagesGetListParams) []types.Message {
	return []types.Message{
		domain.NewMessage("test1"),
		domain.NewMessage("test2"),
	}
}
