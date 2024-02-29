package models

import "time"

type Chat struct {
	ID        int       `json:"id"         db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ChatName  string    `json:"chat_name"  db:"chat_name"`
}

func (c *Chat) InitCreatedAt(){
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}