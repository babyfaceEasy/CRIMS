package models

import "time"

type CloudResource struct {
	ID         uint      `gorm:"primarykey;auto_increment" json:"-"`
	UID        string    `gorm:"<-:false" json:"id,omitempty"`
	Name       string    `gorm:"unique;not null"`
	Type       string    `gorm:"not null"`
	Region     string    `gorm:"not null"`
	CustomerID uint      `gorm:"index"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
