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
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))

	if err != nil{
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("db connection fail:", err)
	}else{
		fmt.Println("Connection established! :)")
	}

	db.MustExec(GetSchema())
	
	chat := models.Chat{}
	err = db.Get(&chat, "SELECT * FROM chats WHERE chat_name=$1", "Public")
	fmt.Println("chat", chat, "err:", err)
	if chat.ChatName == ""{
		tx := db.MustBegin()

		chat.ID = 0
		chat.ChatName = "Public"
		chat.CreatedAt = time.Now()
		tx.NamedExec("INSERT INTO chats (id, created_at, chat_name) VALUES (:id, :created_at, :chat_name)", &chat)
		tx.Commit()
	}
}

func GetDB() *sqlx.DB{
	return db
}