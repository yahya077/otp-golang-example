package model

import (
	auth "github.com/yahya077/otp-golang"
	"time"
)

type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
	Role string `json:"role"`
	auth.OtpModel
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
