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
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"

	"MessengerV2/api/cloudinary"
)

type UserController struct {
	db         *sqlx.DB
	sessionLen int
}

func (u *UserController) Init() {
	u.db = database.GetDB()
	u.sessionLen = 50000 //in seconds, so about 14 hrs
}

func (u *UserController) CheckAuth(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("You're logged in")
}

func (u *UserController) ChangeUserProfilePicture(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil{
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if file.Size > MAX_SIZE{
		return c.Status(http.StatusBadRequest).SendString("File too big.")
	}

	res, err := cloudinary.UploadImage(file)

	fmt.Println("res url:", res.URL)

	return nil
}

func (u *UserController) GetUsername(c *fiber.Ctx) error {
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"username\" field.")
	}

	return c.Status(http.StatusOK).SendString(username)
}

func (u *UserController) GetUserAnonymousStatus(c *fiber.Ctx) error {
	isAnonymous, anonymousErr := c.UserContext().Value("token").(jwt.MapClaims)["is_anonymous"].(bool)

	if !anonymousErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"is_anonymous\" field.")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"is_anonymous": isAnonymous})
}

func (u *UserController) Signin(c *fiber.Ctx) error {
	user := models.User{DB: u.db}
	possibleUser := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	//Check the database to see if the user was there, and compare their password to the user provided one.
	usernameErr := u.db.Get(&possibleUser, "SELECT * FROM users WHERE username=$1", user.Username)
	passwordErr := bcrypt.CompareHashAndPassword([]byte(possibleUser.Password), []byte(user.Password))

	if usernameErr != nil || passwordErr != nil {
		return c.Status(http.StatusNotFound).SendString("Username or Password incorrect. Please try again.")
	}

	user.DisplayName = user.Username
	minifiedUser := struct{
		DisplayName    string `json:"display_name"`
		ProfilePicture string `json:"profile_picture"`
	}{
		DisplayName: user.Username,
		ProfilePicture: user.ProfilePicture.String,
	}
	u.setCookie(c, u.getJwtToken(user), u.sessionLen)

	return c.Status(http.StatusOK).JSON(minifiedUser)
}

func (u *UserController) Signup(c *fiber.Ctx) error {
	user := models.User{DB: u.db}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	//Before validating, trim the password and username for spaces to get the "real" length.
	user.Username = strings.Trim(user.Username, " ")
	user.Password = strings.Trim(user.Password, " ")
	user.DisplayName = user.Username

	//Add password gen for anonymous user to clear password validation
	if user.IsAnonymous {
		randomPassword, _ := password.Generate(12, 8, 0, false, false)
		user.Password = randomPassword
	}

	//Validate for all user credentials.
	if err := user.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.InitCreatedAt()
	user.Password = string(hashedPassword)
	result, _ := u.db.PrepareNamed("INSERT INTO users (created_at, updated_at, username, password, is_anonymous) " +
		"VALUES (:created_at, :updated_at, :username, :password, :is_anonymous) RETURNING id")

	//Execute the prepared statement. This will allow the id of the created user to be returned.
	if err := result.Get(&user.ID, user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	u.setCookie(c, u.getJwtToken(user), u.sessionLen)
	return c.Status(http.StatusCreated).JSON(user)
}

func (u *UserController) Signout(c *fiber.Ctx) error {
	isAnonymous, anonymousErr := c.UserContext().Value("token").(jwt.MapClaims)["is_anonymous"].(bool)

	if !anonymousErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"is_anonymous\" field.")
	}

	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"username\" field.")
	}

	if isAnonymous {
		result, _ := u.db.Exec("DELETE FROM users WHERE username=$1", username)
		rowsAffected, _ := result.RowsAffected()

		fmt.Println("rows:", rowsAffected)
	}

	u.setCookie(c, "", 0)
	return c.Status(http.StatusOK).SendString("signed out")
}

func (u *UserController) GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	if err := u.db.Select(&users, "SELECT username, profile_picture, is_anonymous FROM users"); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (u *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE id=$1", id); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("No user with id %s found.", id))
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	user := models.User{}

	if err := u.db.Get(&user, "SELECT * FROM users WHERE username=$1", username); err != nil {
		fmt.Println("err:", err)
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("No user with username %s found.", username))
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	result, _ := u.db.Exec("DELETE FROM users WHERE id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("No user with id %s found.", id))
	}

	return c.Status(http.StatusOK).SendString(fmt.Sprintf("Deleted user with id \"%s\"", id))
}

func (u *UserController) getJwtToken(user models.User) string {
	expiry := time.Now().Add(time.Duration(u.sessionLen) * time.Second)
	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":          expiry.Unix(),
		"username":     user.Username,
		"is_anonymous": user.IsAnonymous,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))

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
		MaxAge:   sessionLen,
		Value:    value,
	})
}
