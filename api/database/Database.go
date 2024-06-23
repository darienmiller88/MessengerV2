package database

import (
	"fmt"
	"os"
	// "time"
	
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"MessengerV2/api/models"
	"MessengerV2/api/SQLConstants"
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

func InsertMessage(message models.Message) (models.Message, error){
	var result *sqlx.NamedStmt
	var err error

	// If the message was sent to a particular chat, insert the message with that chat id.
	if message.ChatID.Valid {
		result, err = db.PrepareNamed(sqlconstants.INSERT_GROUP_CHAT_MESSAGE)
	}else{
		result, err = db.PrepareNamed(sqlconstants.INSERT_PUBLIC_MESSAGE)
	}

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
	_, err       := db.Exec(sqlconstants.DELETE_USER_FROM_GROUPCHAT, chatId, username)

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sqlx.DB{
	return db
}