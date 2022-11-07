package domain

import "time"

type GameStatus string

const (
	STARTED  GameStatus = "STARTED"
	FINISHED GameStatus = "FINISHED"
)

type Game struct {
	Id        string     `gorm:"primaryKey" json:"id"`
	Status    GameStatus `gorm:"type:game_status"`
	Answers   []string   `gorm:"type:answers"`
	UserId    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
