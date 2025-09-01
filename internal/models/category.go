package models

type Category struct {
	Id       int64     `gorm:"primaryKey;auto_increment" json:"id"`
	Name     string    `gorm:"size:255;not null;index" json:"name"`
	Products []Product `gorm:"foreignkey:CategoryId" json:"products"`
}
