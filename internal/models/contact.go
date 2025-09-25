package models

import "time"

// Contact avec GORM tags pr persistance et export
type Contact struct {
	ID        int       `gorm:"primaryKey; autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100); not null" json:"nom"`
	Email     string    `gorm:"type:varchar(255); not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
