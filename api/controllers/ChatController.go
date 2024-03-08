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

func (c *ChatController) GetPrivateChats(fc *fiber.Ctx) error{
	chats    := []models.Chat{}
	username := fc.Params("username")

	if err := c.db.Select(&chats, "SELECT chats.* FROM chats " +
								    "JOIN user_chats ON chats.id = user_chats.chat_id " +
									"JOIN users ON users.username = user_chats.username " +
									"WHERE users.username=$1", username); err != nil{
		return fc.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	return fc.Status(http.StatusOK).JSON(chats)
}

func (c *ChatController) AddNewChat(fc *fiber.Ctx) error{
	chat          := models.Chat{}
	chatWithUsers := struct{
		ChatName   string `json:"chat_name"`
		PictureUrl string `json:"picture_url"` 
		Users    []string `json:"users"`
	}{}

	if err := fc.BodyParser(&chatWithUsers); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}
	
	//Add the necessary data to the new chat the user will create.
	chat.InitCreatedAt()
	chat.ChatName   = chatWithUsers.ChatName
	chat.PictureUrl = chatWithUsers.PictureUrl
	chat, err      := database.CreateNewChat(chat)
	
	if err != nil {
		return fc.Status(http.StatusInternalServerError).SendString(err.Error())
	}
		
	//Since Users and Chats has a many to many relationship, inserting a new group is a two step process, by first
	//retrieving the id of the newly created Chat, and creating new rows in the join table with the users.
	userChat := models.UserChat{
		ChatID: chat.ID,
	}

	userChat.InitCreatedAt()
	for _, user := range chatWithUsers.Users {
		userChat.Username = user
		userChat, err = database.CreateUserChat(userChat)

		if err != nil{
			return fc.Status(http.StatusNotFound).SendString(err.Error())
		}
	}
	
	return fc.Status(http.StatusOK).JSON(chat)
}

func (c *ChatController) DeleteChat(fc *fiber.Ctx) error{
	id := fc.Params("id")

	result, _       := c.db.Exec("DELETE FROM chats WHERE id=$1", id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fc.Status(http.StatusNotFound).SendString(fmt.Errorf("no chat with id %s found", id).Error())
	}

	return fc.Status(http.StatusOK).SendString("id: " + id)
}