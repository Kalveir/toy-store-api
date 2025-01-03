package models

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Price       uint      `json:"price" gorm:"not null"`
	Stock       uint      `json:"stock" gorm:"default:0"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	Review      []Review  `json:"review"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
