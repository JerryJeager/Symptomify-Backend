package chats

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatStore interface {
	CreateChat(ctx context.Context, chat *Chat) error
	GetChatByTabID(ctx context.Context, tabID uuid.UUID) (*[]Chat, error)
	DeleteChat(ctx context.Context, chatID uuid.UUID) error
}

type ChatRepo struct {
	client *gorm.DB
}

func NewChatRepo(client *gorm.DB) *ChatRepo {
	return &ChatRepo{client: client}
}

func (r *ChatRepo) CreateChat(ctx context.Context, chat *Chat) error {
	return r.client.WithContext(ctx).Create(chat).Error
}

func (r *ChatRepo) GetChatByTabID(ctx context.Context, tabID uuid.UUID) (*[]Chat, error) {
	var chats []Chat
	if err := r.client.WithContext(ctx).Find(&chats, "tab_id = ?", tabID).Error; err != nil {
		return nil, err
	}
	return &chats, nil
}

func (r *ChatRepo) DeleteChat(ctx context.Context, chatID uuid.UUID) error {
	qry := r.client.WithContext(ctx).Delete(&Chat{}, "id = ?", chatID)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("chat with this id does not exist")
	}
	return nil
}
