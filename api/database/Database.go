package database

import (
	"MessengerV2/api/models"
	"fmt"
	"os"
	"time"

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
	
	chat := models.Chat{}
	err = _db.Get(&chat, "SELECT * FROM chats WHERE chat_name=$1", "Public")
	fmt.Println("chat", chat, "err:", err)
	if chat.ChatName == ""{
		tx := _db.MustBegin()

		chat.ID = 0
		chat.ChatName = "Public"
		chat.CreatedAt = time.Now()
		chat.UpdatedAt = time.Now()
		tx.NamedExec("INSERT INTO chats (id, created_at, updated_at, chat_name) " +
		"VALUES (:id, :created_at, :updated_at, :chat_name)", &chat)
		tx.Commit()
	}

	db = _db
}

func GetDB() *sqlx.DB{
	return db
}