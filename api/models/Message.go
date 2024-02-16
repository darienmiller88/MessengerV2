package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"
	// "github.com/nerock/ozzo-validation/is"
)

type Message struct{    
	ID             int            `json:"id"              db:"id"`
	CreatedAt      time.Time      `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"      db:"updated_at"`
	Receiver       sql.NullString `json:"receiver"        db:"receiver"`
	ChatID         sql.NullInt64  `json:"-"               db:"chat_id"`
	ImageURL       sql.NullString `json:"image_url"       db:"image_url"`
	MessageContent string         `json:"message_content" db:"message_content"`
	MessageDate    string         `json:"message_date"    db:"message_date"`
	Username       string         `json:"username"        db:"username"`

	//This model will keep a reference to the database to verify usernames.
	DB             *sqlx.DB		  `json:"-"`
}

func (m *Message) Validate() error{
	return validation.ValidateStruct(
		m, 
		validation.Field(&m.MessageContent, validation.Required),
		validation.Field(&m.Username, validation.Required, validation.By(m.CheckUsername)),
	)
}

//Function to ensure the message sent to the server has a username that is in the database
func (m *Message) CheckUsername(val interface{}) error{
	username, success := val.(string)
	user := User{}

	if !success {
		return errors.New("Error parsing value, expected string.")
	}

	if err := m.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username); err != nil{
		return errors.New(fmt.Sprintf("Username \"%s\" was not found.", username))
	}

	return nil
}

func (m *Message) InitCreatedAt(){
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}