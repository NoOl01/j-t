package models

type Product struct {
	Id          int64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string   `gorm:"size:255;not null" json:"name"`
	Description string   `gorm:"not null" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	CategoryId  int64    `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"foreignkey:CategoryId" json:"category"`
}
