package controllers

import (
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ChatController struct{
	db *sqlx.DB
}

func (c *ChatController) Init(){
	c.db = database.GetDB()
}

func (c *ChatController) GetChats(fc *fiber.Ctx) error{
	chats := []models.Chat{}

	if err := c.db.Select(&chats, "SELECT * FROM chats"); err != nil{
		return fc.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	return fc.Status(http.StatusOK).JSON(chats)
}

func (c *ChatController) AddNewChat(fc *fiber.Ctx) error{
	chat          := models.Chat{}
	// userChat      := models.UserChat{}
	chatWithUsers := struct{
		ChatName string   `json:"chat_name"`
		Users    []string `json:"users"`
	}{}

	if err := fc.BodyParser(&chatWithUsers); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	chat.InitCreatedAt()
	chat.ChatName = chatWithUsers.ChatName
	// chat, err := database.CreateNewChat(chat)

	// if err != nil {
	// 	return fc.Status(http.StatusInternalServerError).SendString(err.Error())
	// }

	fmt.Println("new chat:", chat)

	return fc.Status(http.StatusOK).JSON(chat)
}

func (c *ChatController) DeleteChat(fc *fiber.Ctx) error{
	return nil
}