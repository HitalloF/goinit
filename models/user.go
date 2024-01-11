package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string  `json:"Name"`
	Email        *string `json:"Email"`
	Telefone     uint    `json:"Telefone"`
	Cnpj_aproved bool    `json:"Cnpj_aproved"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
