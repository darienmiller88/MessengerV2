package routes

import(
	"github.com/gofiber/fiber/v2"
	"MessengerV2/api/controllers"
)

type MessagesRoutes struct{
	Router           *fiber.App
	messageController controllers.MessageController
}

func (m *MessagesRoutes) Init(){
	m.Router = fiber.New()
	m.messageController.Init()

	m.Router.Post("/userTyping", m.messageController.UserTyping)
	m.Router.Post("/", m.messageController.PostMessage)
	m.Router.Get("/", m.messageController.GetMessages)
	m.Router.Get("/:id", m.messageController.GetMessage)
	m.Router.Delete("/:id", m.messageController.DeleteMessage)
}