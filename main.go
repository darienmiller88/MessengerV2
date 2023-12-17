package main

import (
	"fmt"
	// "net/http"
	"os"

	// "time"

	"MessengerV2/api/database"
	"MessengerV2/api/pusherclient"
	"MessengerV2/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
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
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, https://messengerv2.netlify.app",
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Mount("/api/v1", index.Router)
	
	fmt.Println("running on port:", os.Getenv("PORT"))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}