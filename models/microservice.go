package models

// Microservice object
type Microservice struct {
	ID     string `gorm:"primary_key"`
	Name   string
	GitURI string
}

// MicroserviceDb object
type MicroserviceDb struct {
	DbModelNoID
	Microservice
}

// TableName returns the tablename for the gorm library
func (m MicroserviceDb) TableName() string {
	return "microservices"
}
