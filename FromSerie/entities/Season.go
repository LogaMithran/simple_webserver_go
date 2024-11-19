package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Season struct {
	gorm.Model

	Id           uuid.UUID `gorm:"primaryKey"`
	Title        string    `gorm:"type:varchar(100);index"`
	SeasonNumber int
	ReleaseDate  time.Time
	Episode      []Episode `gorm:"foreignKey:SeasonNumber"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
