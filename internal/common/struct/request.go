package _struct

type NewDoctorRequest struct {
	Name      string `json:"name" validate:"min = 1"`
	ContactNo string `json:"contact_no" validate:"len = 10,regexp = ^[0-9]+$"`
}

type UpdateDoctorRequest struct {
	ContactNo string `json:"contact_no" validate:"len = 10,regexp = ^[0-9]+$"`
}

type NewPatientRequest struct {
	Name      string `json:"name" validate:"min = 1"`
	ContactNo string `json:"contact_no" validate:"len = 10,regexp = ^[0-9]+$"`
	Address   string `json:"address" validate:"min = 1"`
	DoctorId  string `json:"doctor_id"`
}

type UpdatePatientRequest struct {
	DoctorId  string `json:"doctor_id,omitempty"`
	ContactNo string `json:"contact_no,omitempty" validate:"len = 10,regexp = ^[0-9]+$"`
	Address   string `json:"address,omitempty" validate:"min = 1"`
}
