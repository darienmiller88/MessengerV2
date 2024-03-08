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
		router.Post("/userTyping",        m.messageController.UserTyping)
		router.Post("/upload-image",      m.messageController.UploadImageAsMessage)
		router.Post("/",                  m.messageController.PostMessage)
		router.Get("/message-history",    m.messageController.GetMessageHistory)
		router.Get("/chat-messages/:id",  m.messageController.GetGroupChatMessages)
		router.Get("/",                   m.messageController.GetPublicMessages)
		router.Get("/:id",                m.messageController.GetMessageByID)
		router.Delete("/:id",             m.messageController.DeleteMessage)
	})
}