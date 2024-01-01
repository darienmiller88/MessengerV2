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
		m.Router.Post("/userTyping", m.messageController.UserTyping)
		m.Router.Post("/", m.messageController.PostMessage)
		m.Router.Get("/", m.messageController.GetPublicMessages)
		m.Router.Get("/:id", m.messageController.GetMessageByID)
		m.Router.Delete("/:id", m.messageController.DeleteMessage)
		m.Router.Get("/message-history/:username",  middlewares.ProtectUser, m.messageController.GetMessageHistory)
	})
}