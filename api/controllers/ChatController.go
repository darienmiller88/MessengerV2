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
	// chat          := models.Chat{}
	// userChat      := models.UserChat{}
	chatWithUsers := struct{
		ChatName string   `json:"chat_name"`
		Users    []string `json:"users"`
	}{}

	if err := fc.BodyParser(chatWithUsers); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	fmt.Println("new chat:", chatWithUsers)

	return fc.Status(http.StatusOK).SendString("Success!")
}

func (c *ChatController) DeleteChat(fc *fiber.Ctx) error{
	return nil
}