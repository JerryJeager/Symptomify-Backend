package users

import (
	"context"

	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *User, otp *Otp) error
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *User, otp *Otp) error {
	err := r.client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).WithContext(ctx).Error; err != nil {
			return err
		}
		if err := tx.Create(otp).WithContext(ctx).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
