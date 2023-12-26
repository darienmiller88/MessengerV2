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

	u.Router.Post("/signout",  u.userController.Signout)
	u.Router.Use("/signin",    middlewares.ProtectSignin).Post("/signin", u.userController.Signin)
	u.Router.Use("/signup",    middlewares.ProtectSignin).Post("/signup", u.userController.Signup)
	u.Router.Get("/",          middlewares.Auth, u.userController.GetUsers)
	u.Router.Get("/:id",       middlewares.Auth, u.userController.GetUserByID)
	u.Router.Delete("/:id",    middlewares.Auth, u.userController.DeleteUser)
	u.Router.Get("/checkauth", middlewares.Auth, u.userController.CheckAuth) //Throw away route to check log in status
	u.Router.Get("/username/:username", middlewares.Auth, middlewares.ProtectUser, u.userController.GetUserByUsername)
	// u.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
	// })
}
