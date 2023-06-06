package doctor

import (
	"gorm.io/gorm"
	"hms/internal/common"
	"hms/internal/common/struct"
	"time"
)

type Repo struct {
	DB *gorm.DB
}

// NewRepo : This function initialize the repo of the doctor.
func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

type IRepo interface {
	FetchById(id string) (Doctor, error)
	Create(data *_struct.NewDoctorRequest) (Doctor, error)
	Update(data *_struct.UpdateDoctorRequest, id string) (Doctor, error)
	FetchPatientByDoctorId(id string) (Doctor, error)
}

// FetchById  : This function is retrieves the details of doctor from the database.
func (r *Repo) FetchById(id string) (Doctor, error) {
	var doctor Doctor
	err := r.DB.Model(doctor).Where("id = ?", id).First(&doctor).Error
	return doctor, err
}

// Create : This function creates the details of a new doctor in the database.
func (r *Repo) Create(data *_struct.NewDoctorRequest) (Doctor, error) {
	doctor := Doctor{Id: common.GenerateUUID(), Name: data.Name, ContactNo: data.ContactNo}
	err := r.DB.Create(&doctor).Error
	return doctor, err
}

// Update : This function updates the details of an existing doctor in the database.
func (r *Repo) Update(data *_struct.UpdateDoctorRequest, id string) (Doctor, error) {
	var doctor Doctor
	updateDoctor := Doctor{ContactNo: data.ContactNo, UpdatedAt: time.Now()}
	err := r.DB.Model(&doctor).Where("id = ?", id).Updates(&updateDoctor).First(&doctor).Error
	return doctor, err
}

// FetchPatientByDoctorId : This function fetches the doctor from database with the slice of patients
func (r *Repo) FetchPatientByDoctorId(id string) (Doctor, error) {
	var doctor Doctor
	err := r.DB.Model(doctor).Preload("Patient").Where("id = ?", id).First(&doctor).Error
	return doctor, err
}
