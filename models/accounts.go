package models
import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Modal 
	ID 			int
	Username 	string
	Password 	string
	Mail 		string
	Community   int
}