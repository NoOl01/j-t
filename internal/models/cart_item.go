package models

type CartItem struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	ProductId int64   `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product"`
	UserId    int64   `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CartId    int64   `gorm:"not null" json:"cart_id"`
	Cart      Cart    `gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Count     int     `gorm:"not null" json:"count"`
}
