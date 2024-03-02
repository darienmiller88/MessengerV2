package models

import (
	"time"

)

type UserChat struct {
	ID        int       `json:"id"         db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Username  string    `json:"username"   db:"username"`
	ChatID    int       `json:"chat_id"    db:"chat_id"`
}

func (u *UserChat) InitCreatedAt(){
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}