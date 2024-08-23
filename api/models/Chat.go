package models

import (
	"errors"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"
)

type Chat struct {
	ID         int       `json:"id"         db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	ChatName   string    `json:"chat_name"  db:"chat_name"`
	PictureUrl string    `json:"picture_url" db:"picture_url"`

	//This model will keep a reference to the database to verify usernames.
	DB *sqlx.DB `json:"-"`
}

func (c *Chat) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.ChatName, validation.Length(2, 15), validation.By(c.validateChatName)),
	)
}

func (c *Chat) InitCreatedAt() {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Chat) validateChatName(value interface{}) error {
	chatName, valid := value.(string)

	if !valid {
		return errors.New("error parsing value, expected string")
	}

	chat := Chat{}

	c.DB.Get(&chat, "SELECT * FROM chats WHERE chat_name = $1", chatName)

	return nil
}
