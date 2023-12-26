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

	m.Router.Post("/userTyping", middlewares.Auth, m.messageController.UserTyping)
	m.Router.Post("/", middlewares.Auth, m.messageController.PostMessage)
	m.Router.Delete("/:id", middlewares.Auth, m.messageController.DeleteMessage)
	m.Router.Get("/", middlewares.Auth, m.messageController.GetMessages)
	m.Router.Get("/:id", middlewares.Auth, m.messageController.GetMessageByID)
	m.Router.Get("/message-history/:username", middlewares.Auth, middlewares.ProtectUser, m.messageController.GetMessageHistory)
}