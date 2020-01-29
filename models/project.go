package models

// Project object
type Project struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Roles []string
}

// ProjectDb object
type ProjectDb struct {
	DbModelNoID
	Project
}

// TableName returns the tablename for the gorm library
func (m ProjectDb) TableName() string {
	return "projects"
}
