package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"
	// "github.com/nerock/ozzo-validation/is"
)

type Message struct {
	ID             int            `json:"id"              db:"id"`
	CreatedAt      time.Time      `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"      db:"updated_at"`
	Receiver       sql.NullString `json:"receiver"        db:"receiver"`
	ChatID         sql.NullInt64  `json:"-"               db:"chat_id"`
	ImageURL       sql.NullString `json:"image_url"       db:"image_url"`
	MessageContent string         `json:"message_content" db:"message_content"`
	MessageDate    string         `json:"message_date"    db:"message_date"`
	Username       string         `json:"username"        db:"username"`
	DisplayName    string         `json:"display_name"    db:"display_name"`

	//This model will keep a reference to the database to verify usernames.
	DB *sqlx.DB `json:"-"`
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

	if url.String == ""{
		return nil
	}

	// Make a HEAD request to the URL to request header content.
	resp, err := http.Head(url.String)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Check if the response status code is within the success range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Check the Content-Type header for image formats
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/png") || strings.HasPrefix(contentType, "image/jpeg") {
		return nil
	}

	return fmt.Errorf("%s is not a valid image url", url.String)
}

func (m *Message) InitCreatedAt() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}