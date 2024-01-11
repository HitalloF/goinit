package models

import (
	"gorm.io/gorm"
)

type Vendas struct {
	gorm.Model
	ID         uint    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uint    `gorm:"index" json:"user_id"`
	ProdutoID  uint    `gorm:"index" json:"produto_id"`
	Quantidade uint    `json:"quantidade"`
	Valor      float64 `json:"valor"`
	Data       string  `json:"data"`
}

func MigrateVendas(db *gorm.DB) {
	db.AutoMigrate(&Vendas{})
}
