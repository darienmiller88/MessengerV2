package routes

import(
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

type MessagesRoutes struct{
	Router           *fiber.App
	messageController controllers.MessageController
}

func (m *MessagesRoutes) Init(){
	m.Router = fiber.New()
	m.messageController.Init()

	m.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		router.Post("/:chat-id",               middlewares.AuthCheckUser, m.messageController.PostMessage)
		router.Post("/public",                 m.messageController.PostMessage)
		router.Post("/userTyping/:chat-id",    middlewares.AuthCheckUser, m.messageController.UserTyping)
		router.Post("/userTyping/public",      m.messageController.UserTyping)
		router.Post("/upload-image/:chat-id",  middlewares.AuthCheckUser, m.messageController.UploadImageAsMessage)
		router.Post("/upload-image/public",    m.messageController.UploadImageAsMessage)
		router.Get("/message-history",         m.messageController.GetMessageHistory)
		router.Get("/chat-messages/:id",       m.messageController.GetGroupChatMessages)
		router.Get("/:id",                     m.messageController.GetMessageByID)
		router.Get("/",                        m.messageController.GetPublicMessages)
		router.Delete("/public/:message-id",   m.messageController.DeleteMessage)
		router.Delete("/:chat-id/:message-id", m.messageController.DeleteMessage)
	})
}