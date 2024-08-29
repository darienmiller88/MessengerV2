package database

import (
	"fmt"
	"os"
	"time"


	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"MessengerV2/api/SQLConstants"
	"MessengerV2/api/models"
)

var db *sqlx.DB

func Init(){
	_db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))

	if err != nil{
		panic(err)
	}

	err = _db.Ping()

	if err != nil {
		fmt.Println("db connection fail:", err)
	}else{
		fmt.Println("Connection established! :)")
	}

	_db.MustExec(GetSchema())
	db = _db
}

//Create a new "user_chat" row, which is a many-to-many relation between the "users" table, and "chats" table.
func CreateUserChat(userChat models.UserChat) (models.UserChat, error){
	result, err := db.PrepareNamed(sqlconstants.INSERT_USER_CHAT)

	if err != nil{
		return models.UserChat{}, err
	}

	if err := result.Get(&userChat.ID, userChat); err != nil{
		return models.UserChat{}, err
	}

	return userChat, nil
}

//Create a new group chat, and insert it into the database, in the "chats" table.
func CreateNewChat(chat models.Chat) (models.Chat, error){
	result, err := db.PrepareNamed(sqlconstants.INSERT_GROUP_CHAT)

	if err != nil{
		return models.Chat{}, err
	}

	if err := result.Get(&chat.ID, chat); err != nil{
		return models.Chat{}, err
	}

	return chat, nil
}

//Create a new user for this website, and insert them into the database
func CreateUser(user models.User) (models.User, error){
	result, err := db.PrepareNamed(sqlconstants.INSERT_USER)

	if err != nil {
		return models.User{}, err
	}

	//Execute the prepared statement. This will allow the id of the created user to be returned.
	if err := result.Get(&user.ID, user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

//Insert a message into the database. Handles User to User, User to public chat, and User to Group chat.
func InsertMessage(message models.Message) (models.Message, error){
	var result *sqlx.NamedStmt
	var sqlQuery string
	var err error

	// If the message was sent to a particular chat, insert the message with that chat id.
	// If the messgage was sent to a particular user, insert the message with the receiver's username
	// Otherwise, insert the message into the public chat, with no chat id or receiver username.
	if message.ChatID.Valid {
		sqlQuery = sqlconstants.INSERT_GROUP_CHAT_MESSAGE
	} else if message.Receiver.Valid {
		sqlQuery = sqlconstants.INSERT_USER_TO_USER_MESSAGE
	}else{
		sqlQuery = sqlconstants.INSERT_PUBLIC_MESSAGE
	}
	
	result, err = db.PrepareNamed(sqlQuery)

	if err != nil{
		return models.Message{}, err
	}

	if err := result.Get(&message.ID, message); err != nil{
		return models.Message{}, err
	}

	return message, nil
}

// Function to delete a user from a group chat using their username, and the Id of the groupchat they belonged to.
func DeleteUserFromGroupChat(chatId string, username string) error {
	_, err := db.Exec(sqlconstants.DELETE_USER_FROM_GROUPCHAT, chatId, username)

	if err != nil {
		return err
	}

	return nil
}

// Function to delete a message from the "messages" table.
func DeleteMessage(messageId string, username string) error {
	result, err := db.Exec(sqlconstants.DELETE_MESSAGE, messageId, username)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 || err != nil {
		return err
	}

	return nil
}

//Takes in the name of the group chat, the url for the group chat picture, and the id to update a row in "chats".
func UpdateGroupChatPictureAndName(chatName string, pictureUrl string, chatId string) error {
	_, err := db.Exec(sqlconstants.UPDATE_GROUPCHAT_PICTURE_AND_NAME, chatName, pictureUrl, time.Now(), chatId)

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sqlx.DB{
	return db
}