package routes

import(
	"github.com/gofiber/fiber/v2"
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"
)

type ChatsRoutes struct{
	Router          *fiber.App
	chatsController controllers.ChatController
}

func (c *ChatsRoutes) Init(){
	c.Router = fiber.New()
	c.chatsController.Init()

	c.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		router.Get("/",       c.chatsController.GetChats)
		router.Post("/",      c.chatsController.AddNewChat)
		router.Delete("/:id", c.chatsController.DeleteChat)
	})
}