## API for Hospital Management System
This is An Create, Read, Update API for a Hospital Management System <br />
Api is Written using **GoLang** <br />
Database Used : **Postgres** <br />

### DATABASE SCHEMA

#### Doctor
```
ID                             Unique id - char
created_at                     timestamp
updated_at                     timestamp
name                           varchar
contact_no                     char
```
#### Patient
```
ID                              Unique id - char
created_at                      timestamp
updated_at                      timestamp
name                            varchar
contact_no                      char
address                         varchar
doctor_id                       char - foriegnKey
```

### APIs Endpoints

The API services provided by the system are :
## Doctor
### /doctor/

* `POST` : Create a new doctor record in the database.

### Post Paramaters
```
{
    "name":          ”abcd”,
    "contact_no":    “1234567890”
}
```

#### /doctor/:id
* `GET` : Get the doctor details for a particular doctor id
* `PUT` : Update a doctor details. (only contact number can be updated)
### Put Paramaters
```
{
    "contact_no":    “1234567890”
}
```
## Patient
#### /patient/

* `POST` : Create a new patient record in the database with doctor id present in the doctor database.

### Post Paramaters
```
{
     "name":          ”xyz”,
     "contact_no":    “1234567890”,
     “address”:       “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
     “doctor_id”:     “10001”

}
```

#### /patient/:id
* `GET` : Get the patient details with the patient id
* `PUT` : Update a patient details. (Only contact number, address and doctor id is allowed to be updated for a patient)
### Put Paramaters
```
{
     “doctor_id”: “10002”
}
```

## Additional Feature
#### /fetchPatientByDoctorId/:id

* `GET` : Fetch details of all patients using doctor id.

