package models

type User struct {
	Id        int64      `gorm:"primaryKey" json:"id"`
	Login     string     `gorm:"size:255;not_null" json:"login"`
	Email     string     `gorm:"size:255;not_null;uniqueIndex" json:"email"`
	Password  string     `gorm:"size:255;not_null" json:"-"`
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
