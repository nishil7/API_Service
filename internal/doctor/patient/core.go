package patient

import (
	"gopkg.in/validator.v2"
	"hms/internal/common/struct"
	"log"
)

type Core struct {
	repository IRepo
}

// NewCore : This function initialize the core of patient
func NewCore(repo IRepo) *Core {
	return &Core{
		repository: repo,
	}
}

type ICore interface {
	FetchById(id string) (_struct.PatientResponse, error)
	Create(data *_struct.NewPatientRequest) (_struct.PatientResponse, error)
	Update(data *_struct.UpdatePatientRequest, id string) (_struct.PatientResponse, error)
}

// GenerateResponse : This Function Generate Response in json format from the patient model.
func GenerateResponse(patient Patient) _struct.PatientResponse {
	patientResponse := _struct.PatientResponse{
		Id:        patient.GetId(),
		CreatedAt: patient.GetCreatedAt(),
		UpdatedAt: patient.GetUpdatedAt(),
		Name:      patient.GetName(),
		ContactNo: patient.GetContactNo(),
		Address:   patient.GetAddress(),
		DoctorId:  patient.GetDoctorId(),
	}
	return patientResponse
}

// FetchById : This function provides logging and calls the repo for the required operation.
func (co *Core) FetchById(id string) (_struct.PatientResponse, error) {
	patient, err := co.repository.FetchById(id)
	if err != nil {
		log.Printf(err.Error())
	}
	patientResponse := GenerateResponse(patient)
	return patientResponse, err
}

// Create : This function provides  logging and calls the repo for the required operation.
func (co *Core) Create(data *_struct.NewPatientRequest) (_struct.PatientResponse, error) {
	err := validator.Validate(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.PatientResponse{}, err
	}
	patient, err := co.repository.Create(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.PatientResponse{}, err
	}
	patientResponse := GenerateResponse(patient)
	return patientResponse, err
}

// Update : This function provides logging and calls the repo for the required operation.
func (co *Core) Update(data *_struct.UpdatePatientRequest, id string) (_struct.PatientResponse, error) {
	err := validator.Validate(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.PatientResponse{}, err
	}
	patient, err := co.repository.Update(data, id)
	if err != nil {
		log.Printf(err.Error())
		return _struct.PatientResponse{}, err
	}
	patientResponse := GenerateResponse(patient)
	return patientResponse, err
}
