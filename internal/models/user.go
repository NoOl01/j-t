package models

type User struct {
	Id        int64      `gorm:"primaryKey" json:"id"`
	Login     string     `gorm:"size:255;not null" json:"login"`
	Email     string     `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"cart_items,omitempty"`
	Orders    []Order    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"orders"`
	Points    Points     `gorm:"foreignKey:UserId" json:"points"`
}
