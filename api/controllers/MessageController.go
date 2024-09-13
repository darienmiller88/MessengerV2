package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"

	sqlconstants "MessengerV2/api/SQLConstants"
	"MessengerV2/api/cloudinary"
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
)

const MAX_SIZE int64 = 5 * 1024 * 1024 //Max number of bytes, which is 5mb or 5,248,000 bytes
const public string = "Public"

type MessageController struct {
	db *sqlx.DB
}

func (m *MessageController) Init() {
	m.db = database.GetDB()
}

func (m *MessageController) UploadImageAsMessage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if file.Size > MAX_SIZE {
		return c.Status(http.StatusBadRequest).SendString("File too big.")
	}

	imageURL, err := cloudinary.UploadImage(file)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	message := models.Message{DB: m.db}

	message.InitCreatedAt()
	message.Username = c.FormValue("username")
	message.DisplayName = c.FormValue("display_name")
	message.MessageContent = c.FormValue("message_content")
	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	message.ImageURL.Valid = imageURL != ""
	message.ImageURL.String = imageURL

	if err := message.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	message, err = m.addFieldsToMessage(c, message)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	
	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) PostMessage(c *fiber.Ctx) error {
	message := models.Message{DB: m.db}

	if err := c.BodyParser(&message); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := message.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	message, err := m.addFieldsToMessage(c, message)

	if err != nil{
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	message := models.Message{}

	if err := m.db.Get(&message, "SELECT * FROM messages WHERE id=$1", id); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": fmt.Sprintf("Message with id %s does not exist.", id)})
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageHistory(c *fiber.Ctx) error {
	messages := []models.Message{}
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return c.Status(http.StatusBadRequest).SendString("Could not parse field \"username\" in token.")
	}

	if err := m.db.Select(&messages, sqlconstants.GET_MESSAGE_HISTORY, username); err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetUserToUserMessages(c *fiber.Ctx) error {
	messages := []models.Message{}
	username := c.Params("username")

	if err := m.db.Select(&messages, sqlconstants.GET_ALL_DM_FOR_USER, username); err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetGroupChatMessages(c *fiber.Ctx) error {
	id := c.Params("id")
	messages := []models.Message{}

	if err := m.db.Select(&messages, sqlconstants.GET_GROUP_CHAT_MESSAGES, id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetPublicMessages(c *fiber.Ctx) error {
	messages := []models.Message{}

	if err := m.db.Select(&messages, sqlconstants.GET_PUBLIC_MESSAGES); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) DeleteMessage(c *fiber.Ctx) error {
	messageId := c.Params("messageid")
	chatId    := c.Params("chatid", public)
	message   := struct{
		Message models.Message `json:"message"`
	}{}

	if err := c.BodyParser(&message); err != nil {
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	if err := pusherclient.GetPusherClient().Trigger(chatId, "delete_message", message.Message); err != nil {
		fmt.Println("err broadcasting messages:", err)
	}

	if err := database.DeleteMessage(messageId, message.Message.Username); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(message)
}

// Function to allow clients on the front end to know when someone is typing.
func (m *MessageController) UserTyping(c *fiber.Ctx) error {
	return m.handleUserLeavingAndUserTypingPushNotifications(c, "user_typing")
}

// Function to allow clients on the front end to know when someone left the chat.
func (m *MessageController) UserLeft(c *fiber.Ctx) error {
	return m.handleUserLeavingAndUserTypingPushNotifications(c, "user_left")
}

//Function will return the name of the pusher channel to trigger events to based on which route was it.
func (m *MessageController) addFieldsToMessage(c *fiber.Ctx, message models.Message) (models.Message, error){
	channelName      := ""
	chatId           := c.Params("chatid", public)

	message.InitCreatedAt()
	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	convChatId, err := strconv.Atoi(chatId)

	//The "chatid" URL Param will either be a number, or default to "Public". If the URL param can be 
	//converted to a number, assign it to the messages "ChatID" field.
	if err == nil {
		message.ChatID.Int64 = int64(convChatId)
		message.ChatID.Valid = true
	}

	channelName = chatId

	//After adding the fields, insert the message into the database and trigger the message to the front end.
	message, err = m.insertAndTriggerMessage(channelName, message)

	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}

//Function will insert a message into the database, trigger the pusher event to the front end, and
//return that message with its table id.
func (m *MessageController) insertAndTriggerMessage(channelName string, message models.Message) (models.Message, error){
	if err := pusherclient.GetPusherClient().Trigger(channelName, "message", message); err != nil {
		fmt.Println("err broadcasting messages for event message", ":",  err)
	}

	message, err := database.InsertMessage(message)

	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}

//Function to combine the logic of notifying the user of other users typing and leaving group chats.
func (m *MessageController) handleUserLeavingAndUserTypingPushNotifications(c *fiber.Ctx, pusherEventName string) error{
	chatId := c.Params("chatid", public)
	data := struct {
		Username string `json:"username"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := pusherclient.GetPusherClient().Trigger(chatId, pusherEventName, data.Username); err != nil {
		fmt.Println("err broadcasting messages for event", pusherEventName, ":",  err)
	}

	return c.Status(http.StatusOK).JSON(data)
}