package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"

	"MessengerV2/api/SQLConstants"
	"MessengerV2/api/database"
	"MessengerV2/api/models"
)

type ChatController struct{
	db *sqlx.DB
}

func (c *ChatController) Init(){
	c.db = database.GetDB()
}



func (c *ChatController) AddNewUsersToChat(fc *fiber.Ctx) error{
	chatId    := fc.Params("chatid")
	usernames := []string{}

	if err := fc.BodyParser(&usernames); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	convertedChatID, err := strconv.Atoi(chatId)

	if err != nil {
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userChatsToSend := []models.UserChat{}
	userChat        := models.UserChat{
		ChatID: convertedChatID,
	}
	userChat.InitCreatedAt()

	//Iterate through each username the user from the front end has requested to delete from the group chat.
	for _, username := range usernames{
		userChat.Username = username
		userChat, err    := database.CreateUserChat(userChat)

		if err != nil{
			return fc.Status(http.StatusBadRequest).SendString(err.Error())
		}

		userChatsToSend = append(userChatsToSend, userChat)
	}

	return fc.Status(http.StatusOK).JSON(userChatsToSend)
}

func (c *ChatController) RemoveUsersFromChat(fc *fiber.Ctx) error{
	chatId    := fc.Params("chatid")
	usernames := []string{}

	if err := fc.BodyParser(&usernames); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, username := range usernames {
		if err := database.DeleteUserFromGroupChat(chatId, username); err != nil{
			return fc.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	return fc.Status(http.StatusOK).JSON(fiber.Map{
		"removed": usernames,
	})
}

func (c *ChatController) GetUsersInGroupchats(fc *fiber.Ctx) error{
	chatName      := fc.Params("chatname")
	users         := []models.User{}
	chatName, err := url.QueryUnescape(chatName)

	if err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := c.db.Select(&users, sqlconstants.GET_USERS_IN_GROUP_CHAT, chatName); err != nil{
		return fc.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	usernames := []string{}

	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	return fc.Status(http.StatusOK).JSON(usernames)
}

func (c *ChatController) GetGroupChats(fc *fiber.Ctx) error{
	chats    := []models.Chat{}
	username := fc.Params("username")

	if err := c.db.Select(&chats, sqlconstants.GET_GROUP_CHATS, username); err != nil{
		return fc.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	return fc.Status(http.StatusOK).JSON(chats)
}

func (c *ChatController) AddNewChat(fc *fiber.Ctx) error{
	chatWithUsers := struct{
		ChatName   string `json:"chat_name"`
		PictureUrl string `json:"picture_url"` 
		Users    []string `json:"users"`
	}{}
		
	if err := fc.BodyParser(&chatWithUsers); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}
		
	chat := models.Chat{}

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

	//Afterwards, insert all of the users in the group chat into the user_chats table, with the id of their group chat.
	for _, user := range chatWithUsers.Users {
		userChat.Username = user
		userChat, err = database.CreateUserChat(userChat)

		if err != nil{
			return fc.Status(http.StatusNotFound).SendString(err.Error())
		}
	}
	
	return fc.Status(http.StatusOK).JSON(chat)
}

func (c *ChatController) LeaveGroupChat(fc *fiber.Ctx) error{
	chatId := fc.Params("chatid")
	username, usernameErr := fc.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return fc.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"username\" field.")
	}

	if err := database.DeleteUserFromGroupChat(chatId, username); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return fc.Status(http.StatusOK).SendString(fmt.Sprintf("Removed user: %s", username))
}

func (c *ChatController) DeleteChat(fc *fiber.Ctx) error{
	id := fc.Params("chatid")

	result, _       := c.db.Exec(sqlconstants.DELETE_CHAT, id)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fc.Status(http.StatusNotFound).SendString(fmt.Errorf("no chat with id %s found", id).Error())
	}

	return fc.Status(http.StatusOK).SendString("id: " + id)
}