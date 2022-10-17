package orm

import (
	"gorm.io/gorm"
)

type GormProfile struct {
	gorm.Model
	Id       string
	Username string
	Name     string
}
