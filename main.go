package main

import (
	"fmt"
	"os"

	"MessengerV2/api/database"
	"MessengerV2/api/pusherclient"
	"MessengerV2/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	// "github.com/gofiber/fiber/v2/middleware/limiter"
)

func main(){
	err := godotenv.Load()
	app := fiber.New()
	index := routes.Index{}
	
	if err != nil{
		fmt.Println("err:", err)
	}
	
	database.Init()
	pusherclient.Init()
	index.Init()
	
	app.Use(cors.New())
	app.Use(logger.New())
	app.Mount("/api/v1", index.Router)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world change"})
	})

	fmt.Println("running on port:", os.Getenv("PORT"))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}