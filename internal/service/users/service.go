package users

import (
	"context"
	"log"
	"os"
	"time"
	"errors"

	"github.com/JerryJeager/Symptomify-Backend/utils"
	"github.com/JerryJeager/Symptomify-Backend/utils/emails"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *User) error
	VerifyUser(ctx context.Context, verifyUserReq *VerifyUserReq) error
	Login(ctx context.Context, loginReq *LoginReq) (string, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
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


func (s *UserServ) VerifyUser(ctx context.Context, verifyUserReq *VerifyUserReq) error {
	user, err := s.repo.GetUserByEmail(ctx, verifyUserReq.Email)
	if err != nil{
		return errors.New("no account is registered with this email address")
	}

	otp, err := s.repo.GetUserOtp(ctx, user.ID)
	if err != nil{
		return err
	}

	if otp.Otp != verifyUserReq.Otp {
		return errors.New("wrong otp code used")
	}

	if otp.ExpiresAt.Compare(otp.CreatedAt) == -1 {
		return errors.New("expired otp code")
	}

	if err := s.repo.VerifyUser(ctx, user.ID); err != nil{
		return err
	}
	return nil
}

func (s *UserServ) Login(ctx context.Context, loginReq *LoginReq) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, loginReq.Email)
	if err != nil{
		return "", errors.New("no account is registered with this email address")
	}
	if !user.IsVerified {
		return "", errors.New("only verified users can login")
	}

	if err := VerifyPassword(loginReq.Password, user.Password); err != nil{
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil{
		return "", err
	}

	return token, nil
}

func (s *UserServ) GetUser(ctx context.Context, userID uuid.UUID)(*User, error ){
	return s.repo.GetUser(ctx, userID)
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
