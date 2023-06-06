package boot

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hms/internal/common/struct"
	doctor2 "hms/internal/doctor"
	"hms/internal/doctor/patient"
	"log"
)

// DatabaseInit : Connect to postgres SQL and Creates the table in the database.
// func NewRepo(DB *gorm.DB) {
//
// }

var PatientController *patient.Controller
var DoctorController *doctor2.Controller

// CreateTable : This Function creates the schema in Database
func CreateTable(DB *gorm.DB) {
	err := DB.AutoMigrate(&doctor2.Doctor{}) // TODO create new function
	if err != nil {
		log.Fatal("Failed To create Table Doctor")
	}
	err = DB.AutoMigrate(&patient.Patient{})
	if err != nil {
		log.Fatal("Failed To create Table Patient")
	}
}

// DatabaseInit : This Function passes the Database Configurations and initialize the database
func DatabaseInit() {
	var DB *gorm.DB
	var err error
	dbAuth := _struct.DatabaseConfig{
		User:     "nishil.bansal",
		Database: "hospital",
		Port:     5432,
		Timezone: "Asia/Shanghai",
	}

	dsn := fmt.Sprintf("host=localhost user=%s dbname=%s port=%d sslmode=disable TimeZone=%s", dbAuth.User, dbAuth.Database, dbAuth.Port, dbAuth.Timezone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To connect to HMS Databse")
	}
	CreateTable(DB) //NewRepo(DB)
	InitEntites(DB) // TODO INIT ENTITES : new repo, controller and core
}

// InitEntites : This Function initializes the Repo, Core and Controller of both Doctor & Patient
func InitEntites(DB *gorm.DB) {
	DoctorRepo := doctor2.NewRepo(DB)
	DoctorCore := doctor2.NewCore(DoctorRepo)
	DoctorController = doctor2.NewController(DoctorCore)
	PatientRepo := patient.NewRepo(DB)
	PatientCore := patient.NewCore(PatientRepo)
	PatientController = patient.NewController(PatientCore)
}
