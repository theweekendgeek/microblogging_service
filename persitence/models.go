package persitence

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	TwitterId string
	Username  string
	Name      string
	Tweets    []Tweet
}

type Tweet struct {
	gorm.Model
	Text      string
	ProfileID uint
}
