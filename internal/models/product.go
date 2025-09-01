package models

type Product struct {
	Id          int64    `gorm:"primaryKey;auto_increment" json:"id"`
	Name        string   `gorm:"size:255;not null;index" json:"name"`
	Description string   `gorm:"not null" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	Image       string   `gorm:"not null" json:"image"`
	CategoryId  int64    `gorm:"not null;index" json:"category_id"`
	Category    Category `gorm:"foreignkey:CategoryId" json:"category"`
}
