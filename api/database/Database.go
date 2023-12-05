package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
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
}

func GetDB() *sqlx.DB{
	return db
}