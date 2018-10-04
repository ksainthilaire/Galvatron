package models
import (
	"github.com/jinzhu/gorm"
)

type Ban struct {
	gorm.Model
	ID			int
	AccountID	int
}