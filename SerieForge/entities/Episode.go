package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Episode struct {
	gorm.Model

	Id            uuid.UUID `gorm:"primaryKey"`
	SeasonNumber  int
	EpisodeNumber int
	EpisodeName   string
	ReleaseDate   time.Time
	Runtime       int

	Characters []Character `gorm:"many2many:episode_character"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
}
