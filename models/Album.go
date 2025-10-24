package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Album represents data about a record Album.
type Album struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title  string    `gorm:"type:varchar(100); not null" json:"title"`
	Artist string    `gorm:"type:varchar(100); not null" json:"artist"`
	Price  float64   `json:"price"`
}

// BeforeCreate hook to generate UUID before creating record
func (album *Album) BeforeCreate(tx *gorm.DB) error {
	if album.ID == uuid.Nil {
		album.ID = uuid.New()
	}
	return nil
}
