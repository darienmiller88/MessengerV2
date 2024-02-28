package database

import (
	"MessengerV2/api/models"
	"fmt"
	"os"
	// "time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func GetMessages(){

}

func CreateNewChat(chat models.Chat) (models.Chat, error){
	result, err := db.PrepareNamed("INSERT INTO chat (created_at, updated_at, chat_name) " +
	"VALUES(:created_at, :updated_at, :chat_name) RETURNING id")

	if err != nil{
		return models.Chat{}, err
	}

	if err := result.Get(&chat.ID, chat); err != nil{
		return models.Chat{}, err
	}

	return chat, nil
}

func CreateUser(user models.User) (models.User, error){
	result, err := db.PrepareNamed("INSERT INTO users (created_at, updated_at, username, password, is_anonymous) " +
		"VALUES (:created_at, :updated_at, :username, :password, :is_anonymous) RETURNING id")

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
	result, err := db.PrepareNamed("INSERT INTO messages (message_content, message_date, created_at, updated_at, username, image_url, display_name) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username, :image_url, :display_name) RETURNING id")

	if err != nil{
		return models.Message{}, err
	}

	if err := result.Get(&message.ID, message); err != nil{
		return models.Message{}, err
	}

	return message, nil
}

func GetDB() *sqlx.DB{
	return db
}