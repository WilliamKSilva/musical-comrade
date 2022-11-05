package domain

import "time"

type User struct {
	Id        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Games     []Game
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
