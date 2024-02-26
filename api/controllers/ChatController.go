package controllers

import (
	"MessengerV2/api/database"
	"MessengerV2/api/models"
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
	
	return nil
}

func (c *ChatController) DeleteChat(fc *fiber.Ctx) error{
	return nil
}