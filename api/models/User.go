package models

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmoiron/sqlx"
	"github.com/nerock/ozzo-validation/is"
)

type User struct{
	ID             int            `json:"id"              db:"id"`
	CreatedAt      time.Time      `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"      db:"updated_at"`
	Username       string         `json:"username"        db:"username"`
	Password       string         `json:"password"        db:"password"`
	ProfilePicture sql.NullString `json:"profile_picture" db:"profile_picture"`
	IsAnonymous    bool           `json:"is_anonymous"    db:"is_anonymous"`

	//Hold a reference to the database to check and see if a username is taken.
	DB             *sqlx.DB       `json:"-"`
}

func (u *User) Validate() error{
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Username, 
			validation.Required, 
			validation.Length(4, 15), 
			is.Alphanumeric,
			validation.Match(regexp.MustCompile("[0-9]")).Error("Username must contain at least one number"),
			validation.By(u.CheckUsername),
		),
		validation.Field(&u.Password, 
			validation.Required,
			validation.Length(6, 50),
			validation.Match(regexp.MustCompile("[a-z]")).Error("Password must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("Password must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile("[0-9]")).Error("Password must contain at least one number"),
		),
	)
}

func (u *User) CheckUsername(val interface{}) error{
	username, success := val.(string)
	user := User{}

	if !success {
		return errors.New("Error parsing value, expected string.")
	}

	if err := u.DB.Get(&user, "SELECT * FROM users WHERE username=$1", username); err == nil{
		return errors.New(fmt.Sprintf("Username \"%s\" is taken.", username))
	}

	fmt.Println(username)

	return nil
}

func (u *User) InitCreatedAt(){
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}