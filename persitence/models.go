package persitence

import (
	"gorm.io/gorm"
)

// Profile represents a row in the profiles table
type Profile struct {
	gorm.Model
	TwitterId string
	Username  string
	Name      string
	Tweets    []Tweet
}

// Tweet represents a row in the tweets table
type Tweet struct {
	gorm.Model
	Text      string
	TwitterID string
	ProfileID uint
}
