package domain

import "time"

type gameStatus string

const (
	STARTED  gameStatus = "STARTED"
	FINISHED gameStatus = "FINISHED"
)

type Game struct {
	Id        string     `gorm:"primaryKey" json:"id"`
	Answers   []string   `json:"answers"`
	Status    gameStatus `gorm:"type:game_status"`
	UserId    string     `json:"user_id"`
	User      User       `gorm:"foreignKey:UserId"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
