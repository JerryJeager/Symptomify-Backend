package users

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JerryJeager/Symptomify-Backend/utils"
	"github.com/JerryJeager/Symptomify-Backend/utils/emails"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *User) error
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (s *UserServ) CreateUser(ctx context.Context, user *User) error {
	userID := uuid.New()
	otpID := uuid.New()

	if err := user.HashPassword(); err != nil {
		return err
	}

	user.ID = userID
	user.IsVerified = false
	var otp Otp
	otp.ID = otpID
	otp.UserID = userID
	otp.ExpiresAt = time.Now().Add(time.Hour * 24 * 5) //expires after five days
	otp.Otp = utils.GetOTP()

	if err := s.repo.CreateUser(ctx, user, &otp); err != nil {
		return err
	}

	if err := sendCreateUserEmail(user, otp.Otp); err != nil {
		log.Print(err)
	}

	return nil
}

func sendCreateUserEmail(user *User, otp string) error {

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetAddressHeader("From", email, "Symptomify")
	m.SetHeader("To", user.Email)
	m.SetAddressHeader("Cc", user.Email, "Symptomify")
	m.SetHeader("Subject", "Verify your Email")

	m.SetBody("text/html", emails.CreateUserMail(user.Name, user.Email, otp))

	d := gomail.NewDialer("smtp.gmail.com", 465, emailUsername, emailPassword)

	// Send the email to user
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
