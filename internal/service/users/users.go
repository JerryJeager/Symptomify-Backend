package users

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type Otp struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Otp       string    `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type VerifyUserReq struct {
	Email string `json:'email' binding:"required"`
	Otp   string `json:"otp" binding:"required"`
}
type LoginReq struct {
	Email string `json:'email' binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func (user *User) HashPassword() error {
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	password = html.EscapeString(strings.TrimSpace(password))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


