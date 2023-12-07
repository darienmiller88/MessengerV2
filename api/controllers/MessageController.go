package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pusher/pusher-http-go/v5"

	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
)

type MessageController struct{
	pusherClient pusher.Client
	db           *sqlx.DB
}

func (m *MessageController) Init(){
	m.pusherClient = pusherclient.GetPusherClient()
	m.db           = database.GetDB()
}

func (m *MessageController) PostMessage(c *fiber.Ctx) error{
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil{
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	if err := message.Validate(); err != nil{
		c.Status(http.StatusBadRequest).JSON(err)
	}

	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	err := m.pusherClient.Trigger("public", "public_message", message)

	if err != nil{
		fmt.Println("err broadcasting messages:", err)
	}
	
	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessage(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "One message"})
}

func (m *MessageController) GetMessages(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "All message"})
}

func (m *MessageController) DeleteMessage(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Delete message"})
}