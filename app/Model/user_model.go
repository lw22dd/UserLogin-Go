package Model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null"`
	Email    string `gorm:"varchar(50);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
