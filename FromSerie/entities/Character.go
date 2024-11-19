package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Character struct {
	gorm.Model

	Id   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"index:idx_character_name,unique"`

	Episodes []Episode `gorm:"many2many:character_episode"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
