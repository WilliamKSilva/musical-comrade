package domain

import "time"

type GameStatus string
type GameMusicGenre string

const (
	STARTED  GameStatus = "STARTED"
	FINISHED GameStatus = "FINISHED"
)

const (
	POP  GameMusicGenre = "POP"
	ROCK GameMusicGenre = "ROCK"
)

type Game struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	Genre     GameMusicGenre `gorm:"type:game_music_genre"`
	Answers   []string       `json:"answers"`
	Status    GameStatus     `gorm:"type:game_status"`
	UserId    string         `json:"user_id"`
	User      User           `gorm:"foreignKey:UserId"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
