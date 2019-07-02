package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Wallet uint64
}
