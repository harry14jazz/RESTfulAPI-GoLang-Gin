package structs

import "github.com/jinzhu/gorm"

type Band struct {
	gorm.Model
	BandName  string
	Leader    string
	BirthYear int
}
