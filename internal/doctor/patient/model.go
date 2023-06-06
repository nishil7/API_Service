package patient

import (
	"time"
)

// TODO in patient
type Patient struct {
	Id        string    `json:"id" gorm:"primaryKey;type:char(5)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	ContactNo string    `json:"contact_no" gorm:"size:10"`
	Address   string    `json:"address" gorm:"type:varchar(255)"`
	DoctorId  string    `json:"doctor_id"`
}

// GetId : This function provides the id of the patient.
func (p Patient) GetId() string {
	return p.Id
}

// GetName : This function provides the name of the patient.
func (p Patient) GetName() string {
	return p.Name
}

// GetCreatedAt : This function provides the Created at time of the patient.
func (p Patient) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetUpdatedAt : This function provides the Updated at time of the patient.
func (p Patient) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetContactNo : This function provides with the contact number.
func (p Patient) GetContactNo() string {
	return p.ContactNo
}

// GetAddress : This function provides with the address of the patient.
func (p Patient) GetAddress() string {
	return p.Address
}

// GetDoctorId : This function provides with the doctor id of the patient.
func (p Patient) GetDoctorId() string {
	return p.DoctorId
}
