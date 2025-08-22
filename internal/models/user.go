package models

type User struct {
	Id       int64  `gorm:"primary_key;auto_increment" json:"id"`
	Login    string `gorm:"size:255;not_null" json:"login"`
	Email    string `gorm:"size:255;not_null" json:"email"`
	Password string `gorm:"size:255;not_null" json:"password"`
}
