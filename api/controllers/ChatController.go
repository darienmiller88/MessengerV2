package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"

	"MessengerV2/api/SQLConstants"
	"MessengerV2/api/cloudinary"
	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
)

type ChatController struct{
	db          *sqlx.DB
	// pusherClient pusher.Client
}

func (c *ChatController) Init(){
	c.db           = database.GetDB()
	// c.pusherClient = pusherclient.GetPusherClient()
}

func (c *ChatController) ChangeGroupChatSettings(fc *fiber.Ctx) error {
	groupchatImageFile, _ := fc.FormFile("groupchat_image")

	if groupchatImageFile != nil && groupchatImageFile.Size > MAX_SIZE{
		return fc.Status(http.StatusRequestEntityTooLarge).JSON(fiber.Map{"errFileTooBig": "File size too big."})
	}
	
	chatId          := fc.FormValue("chat_id")
	chatName        := fc.FormValue("chat_name")
	currentImageUrl := fc.FormValue("current_image_url")
	err             := validation.Validate(chatName, 
		validation.Length(2, 15),
		validation.Match(regexp.MustCompile("[a-z]|[A-Z]")).Error("Chat name must contain at least one letter."),
	)

	if err != nil{
		return fc.Status(http.StatusBadRequest).JSON(fiber.Map{"errInvalidName": err.Error()})
	}

	var imageUrl string

	//Upload a file to cloudinary only if the user provided a file. If not, do not upload, and simply update the display name.
	if groupchatImageFile != nil {
		imageUrl, err = cloudinary.UploadImage(groupchatImageFile)
	
		if err != nil {
			return fc.Status(http.StatusInternalServerError).SendString(err.Error())
		}
	}else{
		imageUrl = currentImageUrl
	}
	
	if err := database.UpdateGroupChatPictureAndName(chatName, imageUrl, chatId); err != nil{
		return fc.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return fc.Status(http.StatusOK).SendString(imageUrl)
}

func (c *ChatController) NotifyUserLeavingGroupChat(fc *fiber.Ctx) error{
	chatId := fc.Params("chatid", public)
	data := struct {
		Username string `json:"username"`
	}{}

	if err := fc.BodyParser(&data); err != nil {
		return fc.Status(http.StatusBadRequest).JSON(err)
	}

	if err := pusherclient.GetPusherClient().Trigger(chatId, "user_left", data.Username); err != nil {
		fmt.Println("err broadcasting messages:", err)
	}

	return fc.Status(http.StatusOK).JSON(data)
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
		// ChatName   string `json:"chat_name"`
		// PictureUrl string `json:"picture_url"` 
		Chat       models.Chat `json:"chat"`
		Users    []string      `json:"users"`
	}{}
		
	if err := fc.BodyParser(&chatWithUsers); err != nil{
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := chatWithUsers.Chat.Validate(); err != nil {
		return fc.Status(http.StatusBadRequest).SendString(err.Error())
	}
		
	chatWithUsers.Chat.InitCreatedAt()
	newChat, err := database.CreateNewChat(chatWithUsers.Chat)
	
	if err != nil {
		return fc.Status(http.StatusInternalServerError).SendString(err.Error())
	}
		
	//Since Users and Chats has a many to many relationship, inserting a new group is a two step process, by first
	//retrieving the id of the newly created Chat, and creating new rows in the join table with the users.
	userChat := models.UserChat{
		ChatID: newChat.ID,
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
	
	return fc.Status(http.StatusOK).JSON(newChat)
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