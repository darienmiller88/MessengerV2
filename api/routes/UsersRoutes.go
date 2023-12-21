package routes

import (
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct{
	Router         *fiber.App
	userController controllers.UserController
}

func (u *UserRoutes) Init(){
	u.Router = fiber.New()
	u.userController.Init()

	u.Router.Use(middlewares.ProtectSignin).Post("/signin", u.userController.Signin)
	u.Router.Post("/signup", u.userController.Signup)
	u.Router.Post("/signout", u.userController.Signout)
	u.Router.Use(middlewares.Auth).Get("/:id", u.userController.GetUserByID)
	u.Router.Use(middlewares.Auth).Delete("/:id", u.userController.DeleteUser)
	u.Router.Use(middlewares.Auth).Get("/username/:username", u.userController.GetUserByUsername)
	u.Router.Use(middlewares.Auth).Get("/", u.userController.GetUsers)
}