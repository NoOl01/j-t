package models

type Category struct {
	Id       int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Products []Product `gorm:"foreignkey:CategoryId" json:"products"`
}
