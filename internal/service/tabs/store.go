package tabs

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TabStore interface {
	CreateTab(ctx context.Context, tab *Tab) error
	GetTabs(ctx context.Context, userID uuid.UUID) (*[]Tab, error)
	DeleteTab(ctx context.Context, tabID uuid.UUID) error
}

type TabRepo struct {
	client *gorm.DB
}

func NewTabRepo(client *gorm.DB) *TabRepo {
	return &TabRepo{client: client}
}

func (r *TabRepo) CreateTab(ctx context.Context, tab *Tab) error {
	return r.client.WithContext(ctx).Create(tab).Error
}

func (r *TabRepo) GetTabs(ctx context.Context, userID uuid.UUID) (*[]Tab, error) {
	var tab []Tab

	if err := r.client.WithContext(ctx).Find(&tab, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &tab, nil
}

func (r *TabRepo) DeleteTab(ctx context.Context, tabID uuid.UUID) error {
	qry := r.client.WithContext(ctx).Delete(&Tab{}, "id = ?", tabID)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("tab with this id does not exist")
	}
	return nil
}
