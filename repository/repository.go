package repository

import (
	auth "github.com/yahya077/otp-golang"
	"github.com/yahya077/otp-golang-example/database"
	"github.com/yahya077/otp-golang-example/model"
	"time"
)

type UserRepository struct {
	User model.User
}

// VerifyOtp IUserSubject interfaces function
func (u UserRepository) VerifyOtp(phone string, code string) auth.OtpCheckerResponse {
	var (
		otpCode       auth.OtpCode
		authenticated = false
		registered    = false
	)

	err := database.Postgresql.DB.Where("phone = ? AND code = ?", phone, code).Last(&otpCode).Error

	if err == nil {
		if authenticated = !otpCode.IsExpired(); authenticated {
			registered = u.Registered(phone)
		}
	}

	return auth.OtpCheckerResponse{
		Authenticated: authenticated,
		Registered:    registered,
		Phone:         phone,
		Expiration:    time.Now().Add(time.Hour * 72),
	}
}

func (u UserRepository) Registered(phone string) bool {
	err := database.Postgresql.DB.Where("phone = ?", phone).First(&u.User).Error

	return err == nil
}

// Register IUserSubject interfaces function
func (u UserRepository) Register(parser func(interface{}) error) error {
	_ = parser(&u.User)
	//database.Postgresql.DB.Create(&u.User)
	return nil
}
