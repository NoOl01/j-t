package models

type Cart struct {
	Id        int64      `gorm:"primaryKey" json:"id"`
	UserId    int64      `gorm:"not null" json:"user_id"`
	User      User       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CartItems []CartItem `gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"cart_items"`
	Price     float64    `gorm:"not null" json:"price"`
}
