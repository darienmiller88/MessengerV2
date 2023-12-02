package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pusher/pusher-http-go/v5"

	// "github.com/gofiber/fiber/v2/middleware/limiter"
)

func main(){
	err := godotenv.Load()
	app := fiber.New()

	if err != nil{
		fmt.Println("err:", err)
	}

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world"})
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

