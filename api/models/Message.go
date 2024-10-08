package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"

	// "github.com/nerock/ozzo-validation/is"
	"MessengerV2/api/utilities"
)

type Message struct {
	ID             int            `json:"id"              db:"id"`
	CreatedAt      time.Time      `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"      db:"updated_at"`
	ChatID         sql.NullInt64  `json:"-"               db:"chat_id"`
	ImageURL       sql.NullString `json:"image_url"       db:"image_url"`
	MessageContent string         `json:"message_content" db:"message_content"`
	MessageDate    string         `json:"message_date"    db:"message_date"`
	Username       string         `json:"username"        db:"username"`
	DisplayName    string         `json:"display_name"    db:"display_name"`

	//This model will keep a reference to the database to verify usernames.
	DB *sqlx.DB `json:"-"`

	// This field will allow the profile picture of the user this message belongs to, to be returned in a join.
	Profile_picture sql.NullString `json:"profile_picture" db:"profile_picture"`
}

func (m *Message) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.MessageContent, validation.Required),
		validation.Field(&m.Username, validation.Required, validation.By(m.CheckUsername)),
		validation.Field(&m.ImageURL, validation.By(m.isValidImage)),
		validation.Field(&m.DisplayName, validation.By(m.CheckDisplayname)),
	)
}

// Function to ensure the message sent to the server has a username that is in the database
func (m *Message) CheckUsername(val interface{}) error {
	username, success := val.(string)
	user := User{}

	if !success {
		return errors.New("error parsing value, expected string")
	}

	if err := m.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username); err != nil {
		fmt.Println("err:", err, "user:", user.DisplayName)
		return fmt.Errorf("username \"%s\" was not found", username)
	}

	return nil
}

// Function to ensure the message sent to the server has a username that is in the database
func (m *Message) CheckDisplayname(val interface{}) error {
	displayName, success := val.(string)
	user := User{}

	if !success {
		return errors.New("error parsing value, expected string")
	}

	if displayName == ""{
		return nil
	}

	if err := m.DB.Get(&user, "SELECT * FROM users WHERE display_name=$1", displayName); err != nil {
		return fmt.Errorf("display name \"%s\" was not found", displayName)
	}

	return nil
}


func (m *Message) isValidImage(val interface{}) error {
	url, success := val.(sql.NullString)

	if !success {
		return errors.New("error parsing value, expected string")
	}

	if strings.Trim(url.String, " ") == ""{
		return nil
	}

	if err := utilities.CheckValidImageURL(url.String); err != nil{
		return err
	}

	return nil	
}

func (m *Message) InitCreatedAt() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}