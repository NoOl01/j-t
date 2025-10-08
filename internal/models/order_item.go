package models

type OrderItem struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	ProductId int64   `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product"`
	UserId    int64   `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	OrderId   int64   `gorm:"not null" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Count     int     `gorm:"not null" json:"count"`
}
