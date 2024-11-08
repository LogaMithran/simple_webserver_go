package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Id           uuid.UUID `gorm:"primaryKey"`
	Title        string    `gorm:"index:idx_movie_name,unique"`
	EpisodeId    int
	OpeningCrawl string
	Director     string
	Producer     string
	ReleaseDate  string
	Url          string

	Peoples []*People `gorm:"many2many:movie_peoples"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (movie *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	movie.Id = uuid.New()
	return
}
