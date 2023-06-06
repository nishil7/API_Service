package patient

import (
	"gorm.io/gorm"
	"hms/internal/common"
	"hms/internal/common/struct"
	"time"
)

type Repo struct {
	DB *gorm.DB
}

// NewRepo : This function intitalize the repo of the patient.
func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

type IRepo interface {
	FetchById(id string) (Patient, error)
	Create(data *_struct.NewPatientRequest) (Patient, error)
	Update(data *_struct.UpdatePatientRequest, id string) (Patient, error)
}

// FetchById : This function retrieves the details of patient from the database.
func (r *Repo) FetchById(id string) (Patient, error) {
	var patient Patient
	err := r.DB.Model(&patient).Where("id = ?", id).First(&patient).Error
	return patient, err
}

// Create : This function creates the details of a new doctor in the database.
func (r *Repo) Create(data *_struct.NewPatientRequest) (Patient, error) {
	patient := Patient{Id: common.GenerateUUID(), Name: data.Name, ContactNo: data.ContactNo, Address: data.Address, DoctorId: data.DoctorId}
	err := r.DB.Create(&patient).Error
	return patient, err
}

// Update : This function updates the details of an existing doctor in the database.
func (r Repo) Update(data *_struct.UpdatePatientRequest, id string) (Patient, error) {
	var patient Patient
	updatePatient := Patient{ContactNo: data.ContactNo, Address: data.Address, DoctorId: data.DoctorId, UpdatedAt: time.Now()}
	err := r.DB.Model(&patient).Where("id = ?", id).Updates(&updatePatient).First(&patient).Error
	return patient, err
}
