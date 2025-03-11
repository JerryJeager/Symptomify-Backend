package chats

import (
	"context"

	"github.com/google/uuid"
)

type ChatSv interface {
	CreateChat(ctx context.Context, userID, tabID uuid.UUID, chat *Chat) (string, error)
	GetChatByTabID(ctx context.Context, tabID uuid.UUID) (*[]Chat, error)
	DeleteChat(ctx context.Context, chatID uuid.UUID) error
}

type ChatServ struct {
	repo ChatStore
}

func NewChatService(repo ChatStore) *ChatServ {
	return &ChatServ{repo: repo}
}

func (s *ChatServ) CreateChat(ctx context.Context, userID, tabID uuid.UUID, chat *Chat) (string, error) {
	id := uuid.New()
	chat.ID = id
	chat.UserID = userID 
	chat.TabID = tabID 

	if err := s.repo.CreateChat(ctx, chat); err != nil{
		return "", nil
	}

	return id.String(), nil
}

func (s *ChatServ) GetChatByTabID(ctx context.Context, tabID uuid.UUID) (*[]Chat, error) {
	return s.repo.GetChatByTabID(ctx, tabID)
}

func (s *ChatServ) DeleteChat(ctx context.Context, chatID uuid.UUID) error {
	return s.repo.DeleteChat(ctx, chatID)
}