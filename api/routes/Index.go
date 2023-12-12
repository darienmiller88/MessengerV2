package routes

import(
	"github.com/gofiber/fiber/v2"
	"MessengerV2/api/controllers"
)

type Index struct{
	Router *fiber.App
	userRoutes UserRoutes
	messageRoutes MessagesRoutes
	socketController controllers.SocketController
}

func (i *Index) Init(){
	i.Router = fiber.New()

	i.userRoutes.Init()
	i.messageRoutes.Init()
	i.socketController.Init()

	i.Router.Mount("/users", i.userRoutes.Router)
	i.Router.Mount("/messages", i.messageRoutes.Router)
}