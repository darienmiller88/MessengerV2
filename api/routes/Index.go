package routes

import(
	"github.com/gofiber/fiber/v2"
)

type Index struct{
	Router *fiber.App
	userRoutes UserRoutes
	messageRoutes MessagesRoutes
}

func (i *Index) Init(){
	i.Router = fiber.New()

	i.userRoutes.Init()
	i.messageRoutes.Init()

	i.Router.Mount("/users", i.userRoutes.Router)
	i.Router.Mount("/messages", i.messageRoutes.Router)
}