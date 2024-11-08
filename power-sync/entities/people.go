package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type People struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;"`
	Name      string    `gorm:"index:idx_name,unique"`
	Height    string
	Mass      string
	HairColor string
	SkinColor string
	EyeColor  string
	BirthYear string
	Gender    string
	HomeWorld string
	Species   string
	Url       string

	Movies []*Movie `gorm:"many2many:people_movies"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (people *People) BeforeCreate(tx *gorm.DB) (err error) {
	people.Id = uuid.New()
	return
}
