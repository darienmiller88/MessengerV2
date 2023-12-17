package models

import (
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/nerock/ozzo-validation/is"
)

type User struct{
	ID             int       `json:"id"              db:"id"`
	CreatedAt      time.Time `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"      db:"updated_at"`
	Username       string    `json:"username"        db:"username"`
	Password       string    `json:"password"        db:"password"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
}

func (u *User) Validate() error{
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required, validation.Length(4, 12), is.Alphanumeric),
		validation.Field(&u.Password, 
			validation.Required,
			validation.Length(6, 50),
			validation.Match(regexp.MustCompile("[a-z]")).Error("Password must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("Password must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile("[0-9]")).Error("Password must contain at least one number"),
		),
	)
}




// func trimPasswordCheck(password interface{}) error{
// 	pass, _ := password.(string)

// 	if trimmedPassword := strings.Trim(pass, " "); len(trimmedPassword) < passwordMin || len(trimmedPassword) > passwordMax{
// 		return errors.New(fmt.Sprintf("Password must be between %d and %d characters.", passwordMin, passwordMax))
// 	}

// 	return nil
// }
