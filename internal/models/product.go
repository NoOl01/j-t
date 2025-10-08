package models

type Product struct {
	Id          int64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string   `gorm:"size:255;not null" json:"name"`
	Description string   `gorm:"not null" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	Image       string   `gorm:"not null" json:"image"`
	CategoryId  int64    `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
}
