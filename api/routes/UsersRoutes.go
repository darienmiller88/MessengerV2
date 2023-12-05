package routes

import(
	"github.com/gofiber/fiber/v2"
	"MessengerV2/api/controllers"
)

type UserRoutes struct{
	Router *fiber.App
	userController controllers.UserController
}

func (u *UserRoutes) Init(){
	u.Router = fiber.New()
	u.userController.Init()

	u.Router.Post("/signin", u.userController.Signin)
	u.Router.Post("/signup", u.userController.Signup)
	u.Router.Post("/signout", u.userController.Signout)
	u.Router.Get("/:id", u.userController.GetUser)
	u.Router.Delete("/:id", u.userController.DeleteUsers)
	u.Router.Get("/", u.userController.GetUsers)
}