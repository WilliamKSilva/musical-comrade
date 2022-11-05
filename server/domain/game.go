package domain

import "time"

type gameStatus string
type GameMusicGenre string

const (
	STARTED  gameStatus = "STARTED"
	FINISHED gameStatus = "FINISHED"
)

const (
	POP  GameMusicGenre = "POP"
	ROCK GameMusicGenre = "ROCK"
)

type Game struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	Genre     GameMusicGenre `gorm:"type:game_music_genre"`
	Answers   []string       `json:"answers"`
	Status    gameStatus     `gorm:"type:game_status"`
	UserId    string         `json:"user_id"`
	User      User           `gorm:"foreignKey:UserId"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
