package api

// Team object
type Team struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Location string
}

// TeamDb object
type TeamDb struct {
	DbModelNoID
	Team
}

// TableName returns the tablename for the gorm library
func (m TeamDb) TableName() string {
	return "teams"
}
