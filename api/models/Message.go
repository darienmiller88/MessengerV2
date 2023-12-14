package models

import (
	"errors"
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	// "github.com/nerock/ozzo-validation/is"
)

type Message struct{    
	Receiver       string `json:"receiver"`
	MessageContent string `json:"message_content" db:"message_content"`
	MessageDate    string `json:"message_date"    db:"message_date"`
	Username       string `json:"username"`
}

func (m *Message) Validate() error{
	return validation.ValidateStruct(
		&m, 
		validation.Field(&m.MessageContent, validation.Required),
		// validation.Field(&m.MessageDate, validation.Required),
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