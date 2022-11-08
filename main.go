package main

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/yahya077/otp-golang"
	"github.com/yahya077/otp-golang-example/database"
	"github.com/yahya077/otp-golang-example/repository"
	"github.com/yahya077/otp-golang-example/service/smsProvider"
)

func main() {
	app := fiber.New()
	database.Initialize()

	authApp := auth.New(app, database.Postgresql.DB, auth.Config{})

	//set sms provider for sending otp code to phone number
	authApp.SetSmsProvider(smsProvider.MockedSmsSenderProviderClient{})

	//set your custom user model as a subject for verify, insert and register
	authApp.SetUserRepository(repository.UserRepository{})

	authApp.Initialize()

	app.Get("/test-auth", authApp.Config.AuthMiddleware, func(ctx *fiber.Ctx) error {
		return ctx.JSON("authorized")
	})

	app.Listen(":3000")
}
