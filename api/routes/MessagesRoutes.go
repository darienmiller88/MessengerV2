package routes

import (
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type MessagesRoutes struct{
	Router           *fiber.App
	messageController controllers.MessageController
}

func (m *MessagesRoutes) Init(){
	m.Router = fiber.New()
	m.messageController.Init()

	m.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		router.Get("/message-history",       m.messageController.GetMessageHistory)
		router.Get("/chat-messages/:id",     m.messageController.GetGroupChatMessages)
		router.Get("/:id",                   m.messageController.GetMessageByID)
		router.Get("/",                      m.messageController.GetPublicMessages)

		router.Use(limiter.New()).Route("/", func(router fiber.Router) {
			router.Post("/public",               m.messageController.PostMessage)
			router.Post("/:chatid",              middlewares.AuthCheckUser, m.messageController.PostMessage)
			router.Post("/userTyping/public",    m.messageController.UserTyping)
			router.Post("/userTyping/:chatid",   middlewares.AuthCheckUser, m.messageController.UserTyping)
			router.Post("/userLeft/public",      m.messageController.UserLeft)
			router.Post("/userLeft/:chatid",     middlewares.AuthCheckUser, m.messageController.UserLeft)
			router.Post("/upload-image/public",  m.messageController.UploadImageAsMessage)
			router.Post("/upload-image/:chatid", middlewares.AuthCheckUser, m.messageController.UploadImageAsMessage)
			router.Delete("/public/:messageid",  m.messageController.DeleteMessage)
			router.Delete("/:chatid/:messageid", middlewares.AuthCheckUser, m.messageController.DeleteMessage)
		})
	})
}