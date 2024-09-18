package routes

import (
	"MessengerV2/api/controllers"
	"MessengerV2/api/middlewares"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type UserRoutes struct {
	Router         *fiber.App
	userController controllers.UserController
}

func (u *UserRoutes) Init() {
	u.Router = fiber.New()
	u.userController.Init()

	u.Router.Post("/signout", limiter.New(limiter.Config{Max: 20, Expiration: 1* time.Minute}), middlewares.Auth, u.userController.Signout)
	u.Router.Post("/signin", limiter.New(limiter.Config{
		Max: 10,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).SendString("Please try again in one minute.")
		},
	}), middlewares.ProtectSignin, u.userController.Signin)
	u.Router.Post("/signup", limiter.New(limiter.Config{
		Max: 5,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).SendString("Please try again in one minute.")
		},
	}), middlewares.ProtectSignin, u.userController.Signup)
	u.Router.Use(middlewares.Auth).Route("/", func(router fiber.Router) {
		router.Get("/",            u.userController.GetUsers)
		router.Get("/username",    u.userController.GetUsername)
		router.Get("/checkauth",   u.userController.CheckAuth) //Throw away route to check log in status
		router.Get("/isAnonymous", u.userController.GetUserAnonymousStatus)
		// u.Router.Get("/:id",                u.userController.GetUserByID)
		// u.Router.Delete("/:id",             u.userController.DeleteUser)
		router.Get("/username/:username", middlewares.ProtectUser, u.userController.GetUserByUsername)
		router.Get("/profile-picture/:username", u.userController.GetUserProfilePicture)
		router.Put("/upload-profile-pic",  u.userController.ChangeUserProfilePicture)
	})
}