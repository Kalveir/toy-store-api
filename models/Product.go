package models
import(
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint     `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Price       uint     `json:"price" gorm:"not null"`
	Stock       uint     `json:"stock" gorm:"default:0"`
	CategoryID  uint     `json:"category_id"`
	Category    Category
}
