package doctor

import (
	"gopkg.in/validator.v2"
	"hms/internal/common/struct"
	"log"
)

type Core struct {
	repository IRepo
}

// NewCore : This function initialize the core of doctor.
func NewCore(repo IRepo) *Core {
	return &Core{
		repository: repo,
	}
}

type ICore interface {
	FetchById(id string) (_struct.DoctorResponse, error)
	Create(data *_struct.NewDoctorRequest) (_struct.DoctorResponse, error)
	Update(data *_struct.UpdateDoctorRequest, id string) (_struct.DoctorResponse, error)
	FetchPatientByDoctorId(id string) ([]_struct.PatientResponse, _struct.DoctorResponse, error)
}

// GenerateResponse : This function Response in json format from the doctor model.
func GenerateResponse(doctor Doctor) _struct.DoctorResponse {
	doctorResponse := _struct.DoctorResponse{
		Id:        doctor.GetId(),
		CreatedAt: doctor.GetCreatedAt(),
		UpdatedAt: doctor.GetUpdatedAt(),
		Name:      doctor.GetName(),
		ContactNo: doctor.GetContactNo(),
	}
	return doctorResponse
}

// GeneratePatientResponse : This function Response in json format from the patient model.
func GeneratePatientResponse(doctor Doctor) []_struct.PatientResponse {
	var patientResponse []_struct.PatientResponse
	for i := 0; i < len(doctor.Patient); i++ {
		log.Println("g")
		temp := _struct.PatientResponse{
			Id:        doctor.Patient[i].GetId(),
			CreatedAt: doctor.Patient[i].GetCreatedAt(),
			UpdatedAt: doctor.Patient[i].GetUpdatedAt(),
			Name:      doctor.Patient[i].GetName(),
			ContactNo: doctor.Patient[i].GetContactNo(),
			Address:   doctor.Patient[i].GetAddress(),
			DoctorId:  doctor.Patient[i].GetDoctorId(),
		}
		patientResponse = append(patientResponse, temp)
	}
	return patientResponse
}

// FetchById : This function provides logging and calls the repo for the required operation.
func (co *Core) FetchById(id string) (_struct.DoctorResponse, error) {
	doctor, err := co.repository.FetchById(id)
	if err != nil {
		log.Printf(err.Error())
		return _struct.DoctorResponse{}, err
	}
	doctorResponse := GenerateResponse(doctor)
	return doctorResponse, err
}

// Create : This function provides  logging and calls the repo for the required operation.
func (co *Core) Create(data *_struct.NewDoctorRequest) (_struct.DoctorResponse, error) {
	err := validator.Validate(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.DoctorResponse{}, err
	}
	doctor, err := co.repository.Create(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.DoctorResponse{}, err
	}
	doctorResponse := GenerateResponse(doctor)
	return doctorResponse, err
}

// Update  : This function provides logging and calls the repo for the required operation.
func (co *Core) Update(data *_struct.UpdateDoctorRequest, id string) (_struct.DoctorResponse, error) {
	err := validator.Validate(data)
	if err != nil {
		log.Printf(err.Error())
		return _struct.DoctorResponse{}, err
	}
	doctor, err := co.repository.Update(data, id)
	if err != nil {
		log.Printf(err.Error())
		return _struct.DoctorResponse{}, err
	}
	doctorResponse := GenerateResponse(doctor)
	return doctorResponse, err
}

// FetchPatientByDoctorId : This function provides logging and calls the repo for the required operation.
func (co *Core) FetchPatientByDoctorId(id string) ([]_struct.PatientResponse, _struct.DoctorResponse, error) {
	doctor, err := co.repository.FetchPatientByDoctorId(id)
	if err != nil {
		log.Printf(err.Error())
		return nil, _struct.DoctorResponse{}, err
	}
	doctorResponse := GenerateResponse(doctor)
	patientResponse := GeneratePatientResponse(doctor)
	return patientResponse, doctorResponse, err
}
