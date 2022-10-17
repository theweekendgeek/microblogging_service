package orm

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Id       string
	Username string
	Name     string
}
