package models

import (
	"time"
)

type Doctor struct {
	//Id         uuid.UUID `json:"id" gorm:"size:5;type:uuid;default:gen_random_uuid();primaryKey"`
	Id         int       `json:"id" gorm:"primaryKey,autoIncrement"`
	Created_at time.Time `json:"created_at" gorm:"default:current_timestamp"`
	Updated_at time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Name       string    `json:"name" gorm:"type:varchar(100)"`
	Contact_no string    `json:"contact_no" gorm:"type:char(10)"`
}
