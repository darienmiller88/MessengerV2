package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	// "github.com/nerock/ozzo-validation/is"
)

type Message struct{    
	ID             int       `json:"id"              db:"id"`
	CreatedAt      time.Time `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"      db:"updated_at"`
	Receiver       string    `json:"receiver"        db:"receiver"`
	MessageContent string    `json:"message_content" db:"message_content"`
	MessageDate    string    `json:"message_date"    db:"message_date"`
	Username       string    `json:"username"        db:"username"`
	ChatID         int       `json:"-"               db:"chat_id"`
}

func (m *Message) Validate() error{
	return validation.ValidateStruct(
		&m, 
		validation.Field(&m.MessageContent, validation.Required),
		validation.Field(&m.Username, validation.Required),
		validation.Field(&m.Receiver, validation.Required, validation.By(m.CheckUsername)),
	)
}

func (m *Message) CheckUsername(val interface{}) error{
	username, success := val.(string)

	if !success {
		return errors.New("Error parsing value")
	}
	
	fmt.Println(username)

	return nil
}