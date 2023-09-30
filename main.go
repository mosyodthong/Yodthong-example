package main

import (
	"fmt"
	"log"
	"order/route"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // limit of 20 MB
	})
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initTimeZone()
	route.Router(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("API_PORT")))
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err.Error())
	}
	time.Local = ict
}
