package controllers

import (
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{
	db *sqlx.DB
}

func (u *UserController) Init(){
	u.db = database.GetDB()
}

func (u *UserController) Signin(c *fiber.Ctx) error{
	user         := models.User{}
	possibleUser := models.User{}

	if err := c.BodyParser(&user); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	usernameErr := u.db.Get(&possibleUser, "SELECT * FROM users WHERE username=$1", user.Username)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(possibleUser.Password), []byte(user.Password))

	if usernameErr != nil || passwordErr != nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": "Username or Password incorrect. Please try again."})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "You're signed in!"})
}

func (u *UserController) Signup(c *fiber.Ctx) error{
	user := models.User{DB: u.db}

 	if err := c.BodyParser(&user); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	//Before validating, trim the password and username for spaces to get the "real" length.
	user.Username = strings.Trim(user.Username, " ")
	user.Password = strings.Trim(user.Password, " ")

	if err := user.Validate(); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	user.InitCreatedAt()
	user.Password = string(hashedPassword)
	result, _ := u.db.PrepareNamed("INSERT INTO users (created_at, updated_at, username, password) " + 
	"VALUES (:created_at, :updated_at, :username, :password) RETURNING id")
	
	if err := result.Get(&user.ID, user); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) Signout(c *fiber.Ctx) error{
	return nil
}

func (u *UserController) GetUsers(c *fiber.Ctx) error{
	users := []models.User{}

	if err := u.db.Select(&users, "SELECT * FROM users"); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (u *UserController) GetUserByID(c *fiber.Ctx) error{
	id   := c.Params("id")
	user := models.User{}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE id=$1", id); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with id %s found.", id)})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) GetUserByUsername(c *fiber.Ctx) error{
	username := c.Params("username")
	user     := models.User{}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE username=$1", username); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with username %s found.", username)})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error{
	id := c.Params("id")

	result, _ := u.db.Exec("DELETE FROM users WHERE id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with id %s found.", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": fmt.Sprintf("Deleted user with id %s", id)})
}