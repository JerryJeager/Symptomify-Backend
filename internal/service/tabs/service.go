package tabs

import (
	"context"

	"github.com/google/uuid"
)

type TabSv interface {
	CreateTab(ctx context.Context, userID uuid.UUID) (string, error)
	GetTabs(ctx context.Context, userID uuid.UUID) (*[]Tab, error)
	DeleteTab(ctx context.Context, tabID uuid.UUID) error
}

type TabServ struct {
	repo TabStore
}

func NewTabService(repo TabStore) *TabServ {
	return &TabServ{repo: repo}
}

func (s *TabServ) CreateTab(ctx context.Context, userID uuid.UUID) (string, error) {
	id := uuid.New() 
	var tab Tab 
	tab.ID = id 
	tab.UserID = userID 

	if err := s.repo.CreateTab(ctx, &tab); err != nil{
		return "", err
	}

	return id.String(), nil
}

func (s *TabServ) GetTabs(ctx context.Context, userID uuid.UUID) (*[]Tab, error){
	return s.repo.GetTabs(ctx, userID)
}

func (s *TabServ) DeleteTab(ctx context.Context, tabID uuid.UUID) error{
	return s.repo.DeleteTab(ctx, tabID)
}