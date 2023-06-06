package _struct

import (
	"time"
)

type DoctorResponse struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ContactNo string    `json:"contact_no"`
	//Patient   []patient.Patient `json:"omitempty"`
}

type PatientResponse struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ContactNo string    `json:"contact_no"`
	Address   string    `json:"address"`
	DoctorId  string    `json:"doctor_id"`
}
