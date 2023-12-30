package routes

import (
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	Router         *fiber.App
	userController controllers.UserController
}

func (u *UserRoutes) Init() {
	u.Router = fiber.New()
	u.userController.Init()

	u.Router.Post("/signout", middlewares.Auth,          u.userController.Signout)
	u.Router.Post("/signin",  middlewares.ProtectSignin, u.userController.Signin)
	u.Router.Post("/signup",  middlewares.ProtectSignin, u.userController.Signup)
	u.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		u.Router.Get("/checkauth",          u.userController.CheckAuth) //Throw away route to check log in status
		u.Router.Get("/",                   u.userController.GetUsers)
		// u.Router.Get("/:id",                u.userController.GetUserByID)
		// u.Router.Delete("/:id",             u.userController.DeleteUser)
		u.Router.Get("/username/:username", middlewares.ProtectUser, u.userController.GetUserByUsername)
	})
}
