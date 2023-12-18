package controllers

import (
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserController struct{
	db *sqlx.DB
}

func (u *UserController) Init(){
	u.db = database.GetDB()
}

func (u *UserController) Signin(c *fiber.Ctx) error{


	return nil
}

func (u *UserController) Signup(c *fiber.Ctx) error{
	user := models.User{}

 	if err := c.BodyParser(&user); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := user.Validate(); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) Signout(c *fiber.Ctx) error{
	return nil
}

func (u *UserController) GetUsers(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "All Users"})
}

func (u *UserController) GetUser(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "One User"})
}

func (u *UserController) DeleteUsers(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Delete user"})
}