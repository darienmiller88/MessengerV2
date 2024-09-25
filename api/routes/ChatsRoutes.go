package routes

import (
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

		router.Use(limiter.New()).Route("/", func(router fiber.Router) {
			router.Post("/addNewChat",                 c.chatsController.AddNewChat)
			router.Post("/add-new-users/:chatid",      middlewares.AuthCheckUser, c.chatsController.AddNewUsersToChat)
			router.Post("/notify-user-left/:chatid",   middlewares.AuthCheckUser, c.chatsController.NotifyUserLeavingGroupChat)
			router.Delete("/remove-users/:chatid",     middlewares.AuthCheckUser, c.chatsController.RemoveUsersFromChat)
			router.Delete("/:chatid",                  middlewares.AuthCheckUser, c.chatsController.DeleteChat)
			router.Delete("/leave-group-chat/:chatid", middlewares.AuthCheckUser, c.chatsController.LeaveGroupChat)
			router.Put("/modify-group-chat/:chatid",   middlewares.AuthCheckUser, c.chatsController.ChangeGroupChatSettings)	
		})
	})
}