package controllers

import (
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	db         *sqlx.DB
	sessionLen int
}

func (u *UserController) Init() {
	u.db = database.GetDB()
	u.sessionLen = 5000 //in seconds
}

func (u *UserController) Signin(c *fiber.Ctx) error {
	user := models.User{}
	possibleUser := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	//Check the database to see if the user was there, and compare their password to the user provided one.
	usernameErr := u.db.Get(&possibleUser, "SELECT * FROM users WHERE username=$1", user.Username)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(possibleUser.Password), []byte(user.Password))

	if usernameErr != nil || passwordErr != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": "Username or Password incorrect. Please try again."})
	}

	u.setCookie(c, u.getJwtToken(user), u.sessionLen)
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "You're signed in!"})
}

func (u *UserController) Signup(c *fiber.Ctx) error {
	user := models.User{DB: u.db}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	//Before validating, trim the password and username for spaces to get the "real" length.
	user.Username = strings.Trim(user.Username, " ")
	user.Password = strings.Trim(user.Password, " ")

	//Validate for all user credentials.
	if err := user.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.InitCreatedAt()
	user.Password = string(hashedPassword)
	result, _ := u.db.PrepareNamed("INSERT INTO users (created_at, updated_at, username, password) " +
		"VALUES (:created_at, :updated_at, :username, :password) RETURNING id")

	//Execute the prepared statement. This will allow the id of the created user to be returned.
	if err := result.Get(&user.ID, user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	u.setCookie(c, u.getJwtToken(user), u.sessionLen)
	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) Signout(c *fiber.Ctx) error {
	u.setCookie(c, "", 0)
 	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "signed out"})
}

func (u *UserController) GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	if err := u.db.Select(&users, "SELECT * FROM users"); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	fmt.Println("context:", c.UserContext().Value("username"))
	return c.Status(http.StatusOK).JSON(users)
}

func (u *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE id=$1", id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with id %s found.", id)})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) GetUserByUsername(c *fiber.Ctx) error {
	usernameFromContext := c.UserContext().Value("username").(string)
	usernameFromURL     := c.Params("username")
	user                := models.User{}

	if usernameFromContext != usernameFromURL {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"err": fmt.Sprintf("User %s is not allowed to "+
			"view information from user %s", usernameFromContext, usernameFromURL)})
	}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE username=$1", usernameFromURL); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with username %s found.", usernameFromURL)})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	result, _ := u.db.Exec("DELETE FROM users WHERE id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": fmt.Sprintf("No user with id %s found.", id)})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": fmt.Sprintf("Deleted user with id %s", id)})
}

func (u *UserController) getJwtToken(user models.User) string{
	expiry := time.Now().Add(time.Duration(u.sessionLen) * time.Second)
	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp"     : expiry.Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

	fmt.Println("token:", tokenString)
	return tokenString
}

func (u *UserController) setCookie(c *fiber.Ctx, value string, sessionLen int) {
	c.Cookie(&fiber.Cookie{
		Name:     "user",
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Expires:  time.Now().Add(time.Duration(sessionLen) * time.Second),
		Value:    value,
	})
}
