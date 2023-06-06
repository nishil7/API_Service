package doctor

import (
	"hms/internal/doctor/patient"
	"time"
)

type Doctor struct {
	Id        string            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time         `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time         `json:"updated_at" gorm:"default:current_timestamp"`
	Name      string            `json:"name" gorm:"type:varchar(100)"`
	ContactNo string            `json:"contact_no" gorm:"type:char(10)"`
	Patient   []patient.Patient `gorm:"association_autoupdate:false"`
}

// GetId : This function provides the id of the doctor.
func (d Doctor) GetId() string {
	return d.Id
}

// GetName : This function provides the name of the doctor.
func (d Doctor) GetName() string {
	return d.Name
}

// GetCreatedAt : This function provides the created at time of doctor.
func (d Doctor) GetCreatedAt() time.Time {
	return d.CreatedAt
}

// GetUpdatedAt : This function provides the updated at time of doctor.
func (d Doctor) GetUpdatedAt() time.Time {
	return d.UpdatedAt
}

// GetContactNo : This function provides the contact number of doctor.
func (d Doctor) GetContactNo() string {
	return d.ContactNo
}
