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

/**
*
*/
func (c *ChatsRoutes) Init(){
	c.Router = fiber.New()
	c.chatsController.Init()

	c.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		router.Get("/private-chats/chats/:username", c.chatsController.GetGroupChats)
		router.Get("/private-chats/users/:chatname", c.chatsController.GetUsersInGroupchats)
		router.Post("/add-new-users/:chatid",        middlewares.AuthCheckUser, c.chatsController.AddNewUsersToChat)
		router.Delete("/remove-users/:chatid",       middlewares.AuthCheckUser, c.chatsController.RemoveUsersFromChat)
		router.Delete("/:chatid",                    middlewares.AuthCheckUser, c.chatsController.DeleteChat)
		router.Delete("/leave-group-chat/:chatid",   middlewares.AuthCheckUser, c.chatsController.LeaveGroupChat)
		router.Post("/",                             c.chatsController.AddNewChat)
	})
}