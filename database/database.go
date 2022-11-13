package database

import (
	auth "github.com/yahya077/otp-golang"
	"github.com/yahya077/otp-golang-example/model"
	postgresql "github.com/yahya077/otp-golang-example/service/database"
)

var Postgresql postgresql.Postgresql

func Initialize() {
	Postgresql.Connect()
	Postgresql.AddToMigrate(&auth.OtpCode{})
	Postgresql.AddToMigrate(&model.User{})
	Postgresql.Migrate()
}
