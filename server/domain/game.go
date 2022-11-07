package domain

import "time"

type Genre string
type Status string

const (
	STARTED  Status = "STARTED"
	FINISHED Status = "FINISHED"
)

const (
	POP  Genre = "POP"
	ROCK Genre = "POP"
)

type Game struct {
	Status    Status    `bson:"status,omitempty" json:"status"`
	Genre     Genre     `bson:"genre,omitempty" json:"genre"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}
