package models

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	ID               uint    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Nome             string  `gorm:"not null" json:"nome"`
	Descricao        string  `gorm:"not null" json:"descricao"`
	Categoria        string  `gorm:"not null" json:"categoria"`
	Marca            string  `gorm:"not null" json:"marca"`
	Preco            float64 `gorm:"not null" json:"preco"`
	PrecoPromocional float64 `json:"preco_promocional,omitempty"`
	DescontoCNPJ     float64 `json:"desconto_cnpj,omitempty"`
	PctKg            float64 `json:"pct_kg,omitempty"`
}

func MigrateProduto(db *gorm.DB) {
	db.AutoMigrate(&Produto{})
}
