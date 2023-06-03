package models

import "time"

type Patient struct {
	Id         int       `json:"id" gorm:"primaryKey,autoIncrement"`
	Created_at time.Time `json:"created_at" gorm:"default:current_timestamp"`
	Updated_at time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Name       string    `json:"name" gorm:"type:varchar(100)"`
	Contact_no string    `json:"contact_no" gorm:"size:10"`
	Address    string    `json:"address" gorm:"type:varchar(255)"`
	Doctor_id  int       `json:"doctor_id"`
	Doctor     Doctor    `gorm:"references:Id;foreignKey: Doctor_id"`
}
