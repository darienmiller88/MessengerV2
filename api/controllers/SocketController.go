package controllers

import (
	"fmt"
	"time"

	"MessengerV2/api/database"

	"github.com/ambelovsky/gosf"
	"github.com/jmoiron/sqlx"
)

type SocketController struct{
	db *sqlx.DB
}

func (s *SocketController) Init(){
	s.db = database.GetDB()

	// s.setOnConnect()
	// s.setPublicEndpoint()
	gosf.OnConnect(s.onConnect)
	gosf.Listen("public_endpoint", s.publicEndPointfunc)
}

func (s *SocketController) onConnect(client *gosf.Client, request *gosf.Request) {
	fmt.Println("Client connected.", request)
	client.Join("public_chat")
}

func (s *SocketController) publicEndPointfunc(client *gosf.Client, request *gosf.Request) *gosf.Message {	
	message := new(gosf.Message)
	message.Success = true
	message.Body = request.Message.Body
	message.Body["message_date"] = time.Now().Format("2006-01-02 3:4:5 pm")

	fmt.Println("message:", message.Body)
	client.Broadcast("public_chat", "receive", message)
	return message
}