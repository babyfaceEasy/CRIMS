package models

import (
	"time"

	"gorm.io/gorm"
)

var CustomerVerboseName = "Customer"

type Customer struct {
	ID             uint             `gorm:"primarykey;auto_increment" json:"-"`
	UID            string           `gorm:"<-:false" json:"id,omitempty"`
	Name           string           `gorm:"size:100;not null" json:"name" `
	Email          string           `gorm:"size:100;not null;unique" json:"email"`
	CloudResources []*CloudResource `gorm:"many2many:customer_cloud_resources;" json:"-"`
	CreatedAt      time.Time        `json:"-"`
	UpdatedAt      time.Time        `json:"-"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"-"`
}
