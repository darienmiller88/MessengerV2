package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/pusher/pusher-http-go/v5"

	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
	"MessengerV2/api/cloudinary"
)

const MAX_SIZE int64 = 5 * 1024 * 1024 //Max number of bytes, which is 5mb or 5,248,000 bytes

type MessageController struct{
	pusherClient pusher.Client
	db           *sqlx.DB
}

func (m *MessageController) Init(){
	m.pusherClient = pusherclient.GetPusherClient()
	m.db           = database.GetDB()
}

func (m *MessageController) UploadImageAsMessage(c *fiber.Ctx) error{
	file, err := c.FormFile("file")

	if err != nil{
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if file.Size > MAX_SIZE{
		return c.Status(http.StatusBadRequest).SendString("File too big.")
	}

	res, err := cloudinary.UploadImage(file)
	
	if err != nil{
		return c.SendString(err.Error())
	}

	message := models.Message{DB: m.db}

	message.InitCreatedAt()
	message.Username        = c.FormValue("username")
	message.MessageContent  = c.FormValue("message_content")
	message.MessageDate     = time.Now().Format("2006-01-02 3:4:5 pm")
	message.ImageURL.Valid  = true
	message.ImageURL.String = res.URL

	result, _ := m.db.PrepareNamed("INSERT INTO messages (message_content, message_date, created_at, updated_at, username, image_url) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username, :image_url) RETURNING id")

	if err := result.Get(&message.ID, message); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	fmt.Println("message:", message)

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) PostMessage(c *fiber.Ctx) error{
	message := models.Message{DB: m.db}

	if err := c.BodyParser(&message); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	if err := message.Validate(); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	
	message.InitCreatedAt()
	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	
	result, _ := m.db.PrepareNamed("INSERT INTO messages (message_content, message_date, created_at, updated_at, username) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username) RETURNING id")
	
	if err := m.pusherClient.Trigger("public", "public_message", message); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	if err := result.Get(&message.ID, message); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}
	
	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageByID(c *fiber.Ctx) error{
	id      := c.Params("id")
	message := models.Message{}

	if err := m.db.Get(&message, "SELECT * FROM messages WHERE id=$1", id); err != nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": fmt.Sprintf("Message with id %s does not exist.", id)})
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageHistory(c *fiber.Ctx) error{
	messages              := []models.Message{}
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string) 

	if !usernameErr{
		return c.Status(http.StatusBadRequest).SendString("Could not parse field \"username\" in token.")
	}

	if err := m.db.Select(&messages, "SELECT * FROM messages WHERE username=$1", username); err != nil{
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetPublicMessages(c *fiber.Ctx) error{
	messages := []models.Message{}

	if err := m.db.Select(&messages, "SELECT * FROM messages WHERE receiver IS NULL AND chat_id IS NULL"); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) DeleteMessage(c *fiber.Ctx) error{
	id      := c.Params("id")
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil{
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	fmt.Println("id:", id)
	if err := m.pusherClient.Trigger("public", "delete_public_message", message); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	result, _       := m.db.Exec("DELETE FROM messages WHERE id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("No user with id %s found.", id))
	}

	return c.Status(http.StatusOK).JSON(message)
}

//Function to allow clients on the front end to know when someone is typing.
func (m *MessageController) UserTyping(c *fiber.Ctx) error{
	data := struct{
		Username string `json:"username"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := m.pusherClient.Trigger("public", "user_typing", data.Username); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	return c.Status(http.StatusOK).JSON(data)
}

//Function to check to see a URL leads to a valid image, with a png or jpg/jpeg extension.
func (m *MessageController) isValidImage(url string) (bool, error) {
	// Make a HEAD request to the URL to request header content.
	resp, err := http.Head(url)
	
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	// Check if the response status code is within the success range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Check the Content-Type header for image formats
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/png") || strings.HasPrefix(contentType, "image/jpeg") {
		return true, nil
	}

	return false, nil
}