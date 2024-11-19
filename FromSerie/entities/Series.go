package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Series struct {
	gorm.Model

	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"index:idx_title,unique"`
	Description  string    `json:"description"`
	ReleaseDate  time.Time `json:"release_date"`
	EndDate      time.Time `json:"end_date"`
	SeasonCount  int       `json:"season_count"`
	EpisodeCount int       `json:"episode_count"`
	Status       string    `json:"status"`
	Rating       float64   `json:"rating"`
	Creators     string    `json:"creators"`
	Cast         string    `json:"cast"`

	Season []Season `gorm:"foreignKey:Title;references:Title"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (series *Series) BeforeCreate(tx *gorm.DB) (err error) {
	series.Id = uuid.New()
	return
}
