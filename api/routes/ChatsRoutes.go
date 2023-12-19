package routes

import(
	"github.com/gofiber/fiber/v2"
	"MessengerV2/api/controllers"
)

type ChatsRoutes struct{
	Router          *fiber.App
	chatsController controllers.ChatController
}

func (c *ChatsRoutes) Init(){
	c.Router = fiber.New()
	c.chatsController.Init()

	c.Router.Get("/", c.chatsController.GetChats)
	c.Router.Delete("/:chatname", c.chatsController.DeleteChat)
	c.Router.Post("/", c.chatsController.AddNewChat)
}