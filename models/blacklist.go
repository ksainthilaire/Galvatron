package models
import (
	"github.com/jinzhu/gorm"
)

type Blacklist struct {
	gorm.Model
	ID 	int
	ip 	string
}