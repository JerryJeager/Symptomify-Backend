package users

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *User, otp *Otp) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
	GetUserOtp(ctx context.Context, userID uuid.UUID) (*Otp, error)
	VerifyUser(ctx context.Context, userID uuid.UUID) error
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

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := r.client.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	var user User
	if err := r.client.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserOtp(ctx context.Context, userID uuid.UUID) (*Otp, error) {
	var otp Otp
	if err := r.client.WithContext(ctx).First(&otp, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}


func (r *UserRepo) VerifyUser(ctx context.Context, userID uuid.UUID) error {
	if err := r.client.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("is_verified", true).Error; err != nil{
		return err
	}
	return nil
}